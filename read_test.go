package be

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var recall1 = NewRecallParam().
	SetRecallName("recall1").
	SetRecallType(RecallTypeX2I).
	SetTriggerItems([]string{"1", "2", "3"})

var recall2 = NewRecallParam().
	SetRecallName("recall2").
	SetRecallType(RecallTypeVector).
	SetTriggerItems([]string{"1,1", "2,2", "3,3"})

var emptyNameRecall = NewRecallParam().
	SetRecallName("").
	SetRecallType(RecallTypeVector).
	SetTriggerItems([]string{"1,1", "2,2", "3,3"})

func CheckRequestNotValidate(t *testing.T, req *ReadRequest) {
	err := req.Validate()
	assert.NotNil(t, err)
	fmt.Printf("Error occur:%s\n", err)
}

func CheckRequestValidate(t *testing.T, req *ReadRequest) {
	err := req.Validate()
	assert.Nil(t, err)
}

func TestReadRequest_BuildUri(t *testing.T) {
	request := NewReadRequest("testBiz", 10).
		AddRecallParam(recall1).
		AddRecallParam(recall2)
	CheckRequestValidate(t, request)

	uri := request.BuildUri()
	got := uri.RequestURI()
	want := "be?biz_name=searcher&recall1_trigger_list=1,2,3&recall1_return_count=0&recall2_return_count=0&p=testBiz&s=testBiz&return_count=10&outfmt=json2&recall2_trigger_list=1,1;2,2;3,3"
	assert.Equal(t, len(got), len(want))
}

func TestReadRequest_Validate(t *testing.T) {
	request := NewReadRequest("testBiz", 10)
	CheckRequestNotValidate(t, request)

	request.AddRecallParam(recall1).AddRecallParam(recall2)
	CheckRequestValidate(t, request)

	request.AddRecallParam(recall1)
	CheckRequestNotValidate(t, request)

	request.SetRecallParams([]RecallParam{}).AddRecallParam(emptyNameRecall)
	CheckRequestValidate(t, request)

	request.AddRecallParam(emptyNameRecall)
	CheckRequestNotValidate(t, request)
}
