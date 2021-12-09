package be

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleFilter_GetConditionValue(t *testing.T) {
	filter := SingleFilter{
		Left:     "key",
		Operator: EQ,
		Right:    "value",
	}
	got := filter.GetConditionValue()
	want := "key==value"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMultiFilter_GetConditionValue(t *testing.T) {
	filter1 := SingleFilter{
		Left:     "key1",
		Operator: EQ,
		Right:    "value",
	}

	filter2 := SingleFilter{
		Left:     "key2",
		Operator: NE,
		Right:    "value",
	}

	multiFilter := MultiFilter{
		Filters:   []Filter{&filter1, &filter2},
		Connector: FilterConnectorAnd,
	}

	got := multiFilter.GetConditionValue()
	fmt.Println(got)
	want := "key1==value AND key2!=value"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
