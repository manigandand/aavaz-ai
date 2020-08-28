package types

import "aavaz/schema"

type AnalysisRes struct {
	Feedbacks []*schema.Analysis `json:"feedback"`
	Analytics *Sentiment         `json:"analytics"`
}

type Sentiment struct {
	Total        int     `json:"total"`
	NetSentiment int     `json:"net_sentiment"`
	Positive     float64 `json:"positive"`
	Negative     float64 `json:"negative"`
	Neutral      float64 `json:"neutral"`
}
