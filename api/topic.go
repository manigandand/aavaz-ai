package api

import (
	"aavaz/errors"
	"aavaz/respond"
	"aavaz/schema"
	"aavaz/types"
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
	// page := respond.NewPage(r)
	analysis, err := store.Analysis().Get()
	if err != nil {
		return err
	}

	respond.OK(w, getStates(analysis))
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

	// TODO: make sure about the length, offset:limit not exceeds length
	return &types.AnalysisRes{
		Feedbacks: analysis[0:10],
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
