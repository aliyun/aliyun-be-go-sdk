package be

type RecallType string

const (
	RecallTypeX2i RecallType = "X2i"
	RecallTypeVector RecallType = "Vector"
)


type FilterClause struct {
	Filter Filter
	Clause string
}

type ScorerClause struct {
	Clause string
}

type RecallParam struct {
	RecallName string
	TriggerItems []string
	RecallType RecallType
	FilterClause FilterClause
	ScorerClause ScorerClause
}

type ReadRequest struct {
	ReturnCount int
	BizName string
	RecallParams []RecallParam
}
