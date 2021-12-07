package be

import (
	"reflect"
	"strings"
)

type Filter interface {
	GetConditionValue() string
}

type FilterOperator string

const (
	EQ FilterOperator = "=="
	NE FilterOperator = "!="
	LT FilterOperator = "<"
	GT FilterOperator = ">"
	LE FilterOperator = "<="
	GE FilterOperator = ">="
)

type FilterConnector string

const (
	FilterConnectorAnd FilterConnector = "AND"
)

type SingleFilter struct {
	Left     string
	Operator FilterOperator
	Right    string
}

type MultiFilter struct {
	Filters   []Filter
	Connector FilterConnector
}

func (f *SingleFilter) GetConditionValue() string {
	return f.Left + string(f.Operator) + f.Right
}

func (f *MultiFilter) GetConditionValue() (string, error) {
	if f.Filters == nil || len(f.Filters) == 0 {
		return "", nil
	}

	var conditions []string
	for _, filter := range f.Filters {
		if filter == nil {
			continue
		}
		filterType := reflect.TypeOf(filter)
		if filterType == reflect.TypeOf(new(MultiFilter)) {
			conditions = append(conditions, "(" + filter.GetConditionValue() + ")")
		} else if filterType == reflect.TypeOf(new(SingleFilter)) {
			conditions = append(conditions, filter.GetConditionValue())
		} else {
			return "", InvalidParamsError{"Illegal filter type:" + filterType.Name()}
		}

	}
	return strings.Join(conditions[:], " " + string(f.Connector) + " "), nil
}
