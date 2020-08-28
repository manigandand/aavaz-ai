package schema

// Analysis struct holds the topic feedback data
type Analysis struct {
	Text   string  `json:"text"`
	Themes []Theme `json:"themes"`
	Score  float64 `json:"score"`
	Date   string  `json:"date"`
}

type Theme struct {
	Topic     string `json:"topic"`
	Theme     string `json:"theme"`
	Sentiment string `json:"sentiment"`
}
