package be

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, got, want)
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
	assert.Equal(t, got, want)
}
