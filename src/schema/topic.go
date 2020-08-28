package schema

// Topic struct holds the overall informations about the topic
type Topic struct {
	Topic     string    `json:"topic"`
	Mentions  int       `json:"mentions"`
	Sentiment Sentiment `json:"sentiment"`
	// 1:Many
	Theme string `json:"theme"`
}

type Sentiment struct {
	Negative string `json:"negative"`
	Positive string `json:"positive"`
	Neutral  string `json:"newtral"`
}
