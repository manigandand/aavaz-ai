package adapter

import (
	"aavaz/errors"
	"aavaz/schema"
)

type Store interface {
	Topic() Topic
	Analysis() Analysis
}

type Topic interface {
	All() ([]*schema.Topic, *errors.AppError)
}

type Analysis interface {
	Get(topics []string) ([]*schema.Analysis, *errors.AppError)
	Search(topic string) ([]*schema.Analysis, *errors.AppError)
}
