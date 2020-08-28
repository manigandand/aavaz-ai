package inmemory

// Analysis implements Analysis adapter interface
type Analysis struct {
	*Client
}

// NewAnalysisStore ...
func NewAnalysisStore(client *Client) *Analysis {
	return &Analysis{client}
}

// func (u *Analysis) tableName() string {
// 	return "abilities"
// }
