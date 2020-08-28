package inmemory

import (
	"aavaz/schema"
	"aavaz/store/adapter"
	"encoding/json"
	"io/ioutil"
	"log"
)

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter() adapter.Store {
	// Load Data
	c := &Client{}
	c.TopicConn = NewTopicStore(c)
	c.AnalysisConn = NewAnalysisStore(c)

	c.loadTopics()
	c.loadTopicsAnalysis()

	return c
}

func (c *Client) loadTopics() {
	var data []*schema.Topic

	file, err := ioutil.ReadFile("./data/themes.json")
	if err != nil {
		log.Fatal("couldn't able to load topics data. " + err.Error())
	}
	if err = json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal("couldn't able to unmarshal topics data. " + err.Error())
	}
	c.topics = data
}

func (c *Client) loadTopicsAnalysis() {
	var data []*schema.Analysis

	file, err := ioutil.ReadFile("./data/f.json")
	if err != nil {
		log.Fatal("couldn't able to load topics data. " + err.Error())
	}
	if err = json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal("couldn't able to unmarshal topics data. " + err.Error())
	}

	// load map
	topicAnalysisMap := make(map[string][]*schema.Analysis)
	for _, ta := range data {
		for _, t := range ta.Themes {
			analysis, ok := topicAnalysisMap[t.Topic]
			if !ok {
				topicAnalysisMap[t.Topic] = []*schema.Analysis{ta}
				continue
			}
			analysis = append(analysis, ta)
			topicAnalysisMap[t.Topic] = analysis
		}
	}

	c.topicAnalysis = data
	c.topicAnalysisMap = topicAnalysisMap
}
