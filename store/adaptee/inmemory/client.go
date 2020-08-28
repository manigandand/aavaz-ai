package inmemory

import (
	"aavaz/schema"
	"aavaz/store/adapter"
)

// Client struct implements the store adapter interface
type Client struct {
	topics []*schema.Topic
	// topicsMap        map[string]*schema.Topic
	topicAnalysis    []*schema.Analysis
	topicAnalysisMap map[string][]*schema.Analysis
	TopicConn        adapter.Topic
	AnalysisConn     adapter.Analysis
}

// Topic ...
func (c *Client) Topic() adapter.Topic {
	return c.TopicConn
}

// Analysis ...
func (c *Client) Analysis() adapter.Analysis {
	return c.AnalysisConn
}
