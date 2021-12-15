package be

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type RecallType string

const (
	RecallTypeX2i    RecallType = "X2i"
	RecallTypeVector RecallType = "Vector"
)

type FilterClause struct {
	Filter Filter
	Clause string
}

func NewFilterClause(filter Filter) *FilterClause {
	return &FilterClause{Filter: filter}
}

func (c *FilterClause) GetFilter() Filter {
	return c.Filter
}

func (c *FilterClause) SetFilter(filter *Filter) *FilterClause {
	c.Filter = *filter
	return c
}

func (c *FilterClause) SetClause(clause string) *FilterClause {
	c.Clause = clause
	return c
}

func (c *FilterClause) BuildParams() string {
	queryClause := ""
	if c.Clause != "" {
		queryClause = c.Clause
	} else {
		queryClause = c.Filter.GetConditionValue()
	}
	return url.QueryEscape(queryClause)
}

type ScorerClause struct {
	Clause string
}

func NewScorerClause(clause string) *ScorerClause {
	return &ScorerClause{Clause: clause}
}

type RecallParam struct {
	RecallName   string
	TriggerItems []string
	RecallType   RecallType
	ScorerClause *ScorerClause
}

func NewRecallParam() *RecallParam {
	return &RecallParam{}
}

func (p *RecallParam) SetRecallName(name string) *RecallParam {
	p.RecallName = name
	return p
}

func (p *RecallParam) SetTriggerItems(items []string) *RecallParam {
	p.TriggerItems = items
	return p
}

func (p *RecallParam) SetRecallType(recallType RecallType) *RecallParam {
	p.RecallType = recallType
	return p
}

func (p *RecallParam) SetScorerClause(clause *ScorerClause) *RecallParam {
	p.ScorerClause = clause
	return p
}

func (p *RecallParam) flatTriggers() string {
	if p.RecallType == RecallTypeX2i {
		return strings.Join(p.TriggerItems[:], ",")
	} else {
		return strings.Join(p.TriggerItems[:], ";")
	}
}

func (p *RecallParam) getTriggerKey() string {
	if p.RecallName == "" {
		return "trigger_list"
	} else {
		return p.RecallName + "_trigger_list"
	}
}

func (p *RecallParam) getScorerKey() string {
	if p.RecallName == "" {
		return "scorer_rule"
	} else {
		return p.RecallName + "_scorer_rule"
	}
}

type ReadRequest struct {
	ReturnCount  int
	BizName      string
	FilterClause *FilterClause
	RecallParams []RecallParam
	QueryParams  map[string]string
}

func NewReadRequest(bizName string, returnCount int) *ReadRequest {
	return &ReadRequest{ReturnCount: returnCount, BizName: bizName, QueryParams: map[string]string{}}
}

func (r *ReadRequest) Validate() error {
	if r.ReturnCount <= 0 {
		return InvalidParamsError{"Return count should be greater than 0"}
	}
	if r.BizName == "" {
		return InvalidParamsError{"Empty biz name"}
	}
	if len(r.RecallParams) == 0 {
		return InvalidParamsError{"Empty recall params"}
	}
	recallNames := map[string]bool{}
	for _, param := range r.RecallParams {
		if recallNames[param.RecallName] {
			return InvalidParamsError{fmt.Sprintf("Duplicate recall name[%s] in RecallParams", param.RecallName)}
		}
		recallNames[param.RecallName] = true
	}
	return nil
}

func (r *ReadRequest) SetReturnCount(returnCount int) *ReadRequest {
	r.ReturnCount = returnCount
	return r
}

func (r *ReadRequest) SetBizName(bizName string) *ReadRequest {
	r.BizName = bizName
	return r
}

func (r *ReadRequest) SetFilterClause(clause *FilterClause) *ReadRequest {
	r.FilterClause = clause
	return r
}

func (r *ReadRequest) SetRecallParams(recallParams []RecallParam) *ReadRequest {
	r.RecallParams = recallParams
	return r
}

func (r *ReadRequest) AddRecallParam(param *RecallParam) *ReadRequest {
	r.RecallParams = append(r.RecallParams, *param)
	return r
}

func (r *ReadRequest) AddQueryParam(key string, value string) *ReadRequest {
	r.QueryParams[key] = value
	return r
}

func (r *ReadRequest) SetQueryParams(params map[string]string) *ReadRequest {
	r.QueryParams = params
	return r
}

func (r *ReadRequest) BuildUri() url.URL {
	uri := url.URL{Path: "be"}

	query := map[string]string{}
	query["biz_name"] = "searcher"
	query["p"] = r.BizName
	query["s"] = r.BizName
	query["return_count"] = strconv.Itoa(r.ReturnCount)
	query["outfmt"] = "json2"

	if r.FilterClause != nil && r.FilterClause.BuildParams() != "" {
		query["filter_rule"] = r.FilterClause.BuildParams()
	}
	for _, recallParam := range r.RecallParams {
		query[recallParam.getTriggerKey()] = recallParam.flatTriggers()
		if recallParam.ScorerClause != nil {
			query[recallParam.getScorerKey()] = recallParam.ScorerClause.Clause
		}
	}

	if len(r.QueryParams) != 0 {
		for k, v := range r.QueryParams {
			query[k] = v
		}
	}

	var params []string
	for k, v := range query {
		params = append(params, k+"="+v)
	}
	uri.RawQuery = strings.Join(params[:], "&")
	return uri
}
