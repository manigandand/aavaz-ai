package inmemory

import (
	"aavaz/errors"
	"aavaz/schema"
)

// Analysis implements Analysis adapter interface
type Analysis struct {
	*Client
}

// NewAnalysisStore ...
func NewAnalysisStore(client *Client) *Analysis {
	return &Analysis{client}
}

// func (a *Analysis) tableName() string {
// 	return "analysis"
// }

func (a *Analysis) Get() ([]*schema.Analysis, *errors.AppError) {
	if len(a.topicAnalysis) == 0 || a.topicAnalysisMap == nil {
		return nil, errors.InternalServer("couldn't able to load the data")
	}

	return a.topicAnalysis, nil
}
