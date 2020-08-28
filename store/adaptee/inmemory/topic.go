package inmemory

import (
	"aavaz/errors"
	"aavaz/schema"
)

// Topic implements Topic adapter interface
type Topic struct {
	*Client
}

// NewTopicStore ...
func NewTopicStore(client *Client) *Topic {
	return &Topic{client}
}

// func (t *Topic) tableName() string {
// 	return "topics"
// }

func (t *Topic) All() ([]*schema.Topic, *errors.AppError) {
	if len(t.topics) == 0 {
		return t.topics, errors.InternalServer("couldn't able to load the data")
	}

	return t.topics, nil
}
