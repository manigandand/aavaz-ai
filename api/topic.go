package api

import (
	"aavaz/errors"
	"aavaz/respond"
	"aavaz/schema"
	"aavaz/types"
	"fmt"
	"math"
	"net/http"
)

func getAllTopicsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	topics, err := store.Topic().All()
	if err != nil {
		return err
	}

	respond.OK(w, topics)
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
	// make sure about the length, offset:limit not exceeds length
	sliceLimit := page.Limit
	if page.Offset+page.Limit >= totalRes {
		sliceLimit = (page.Offset + page.Limit) - totalRes
	}
	if page.Offset < totalRes {
		fmt.Println("less offset", page.Offset, " ", sliceLimit)
		res.Feedbacks = analysis[page.Offset:(page.Offset + sliceLimit)]
	}
	if page.Offset >= totalRes { // no result
		fmt.Println("greater offset", page.Offset, " ", sliceLimit)
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
