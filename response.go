package be

type Response struct {
	Result Result `json:"result"`
}

func NewResponse(result Result) *Response {
	return &Response{Result: result}
}

type MatchItem struct {
	FieldNames  []string        `json:"field_names"`
	FieldValues [][]interface{} `json:"field_values"`
}

type ReadResult struct {
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	MatchItems   MatchItem `json:"match_items"`
}

type WriteResult struct {
	Errno int `json:"errno"`
}

type Result struct {
	MatchItems *MatchItem `json:"match_items"`
}
