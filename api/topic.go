package api

import (
	"aavaz/errors"
	"aavaz/respond"
	"aavaz/schema"
	"aavaz/types"
	"math"
	"net/http"
	"sort"
)

func getAllTopicsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	topics, err := store.Topic().All()
	if err != nil {
		return err
	}

	respond.OK(w, topics)
	return nil
}

func searchTopicsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	if err := r.ParseForm(); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	topic := r.URL.Query().Get("topic")
	analysis, err := store.Analysis().Search(topic)
	if err != nil {
		return err
	}

	respond.OK(w, analysis)
	return nil
}

func getTopicAnalysisHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	if err := r.ParseForm(); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}
	page := respond.NewPage(r)
	analysis, err := store.Analysis().Get(page.Topics)
	if err != nil {
		return err
	}

	res := getStates(analysis)
	totalRes := len(analysis)
	if page.Sort == "date" {
		sort.Sort(SortByDate(analysis))
	}

	// make sure about the length, offset:limit not exceeds length
	sliceLimit := page.Limit
	if page.Offset+page.Limit >= totalRes {
		sliceLimit = (page.Offset + page.Limit) - totalRes
	}
	if page.Offset < totalRes {
		res.Feedbacks = analysis[page.Offset:(page.Offset + sliceLimit)]
	}
	if page.Offset >= totalRes { // no result
		res.Feedbacks = make([]*schema.Analysis, 0)
	}

	respond.Paginate(w, r, res, page, totalRes, len(res.Feedbacks))
	return nil
}

func getStates(analysis []*schema.Analysis) *types.AnalysisRes {
	var total, positive, negative, neutral int
	for _, ays := range analysis {
		for _, t := range ays.Themes {
			total++
			switch t.Sentiment {
			case "negative":
				negative++
			case "positive":
				positive++
			default:
				neutral++
			}
		}
	}

	return &types.AnalysisRes{
		Analytics: &types.Sentiment{
			Total:        total,
			NetSentiment: positive - negative,
			Positive:     PercentOf(positive, total),
			Negative:     PercentOf(negative, total),
			Neutral:      PercentOf(neutral, total),
		},
	}
}

func PercentOf(part int, total int) float64 {
	return math.Floor((float64(part) * float64(100)) / float64(total))
}

type SortByDate []*schema.Analysis

func (a SortByDate) Len() int           { return len(a) }
func (a SortByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }
