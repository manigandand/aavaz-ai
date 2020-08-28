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

func (a *Analysis) Get(topics []string) ([]*schema.Analysis, *errors.AppError) {
	if len(a.topicAnalysis) == 0 || a.topicAnalysisMap == nil {
		return nil, errors.InternalServer("couldn't able to load the data")
	}
	if len(topics) == 0 {
		return a.topicAnalysis, nil
	}

	visted := make(map[string]bool)
	res := make([]*schema.Analysis, 0)
	for _, t := range topics {
		if _, ok := visted[t]; !ok {
			visted[t] = true
			if ays, ok := a.topicAnalysisMap[t]; ok {
				res = append(res, ays...)
			}
		}
	}

	return res, nil
}
