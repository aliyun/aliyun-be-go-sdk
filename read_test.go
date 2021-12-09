package be

import (
	"fmt"
	"reflect"
	"testing"
)

var recall1 = NewRecallParam().
	SetRecallName("recall1").
	SetRecallType(RecallTypeX2i).
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
	if err == nil {
		t.Errorf("Error should not empty")
	}
	fmt.Printf("Error occur:%s\n", err)
}

func CheckRequestValidate(t *testing.T, req *ReadRequest) {
	err := req.Validate()
	if err != nil {
		t.Errorf("Error should be empty, error:%s", err)
	}
}

func TestReadRequest_BuildUri(t *testing.T) {
	// request
	request := NewReadRequest("testBiz", 10).
		AddRecallParam(recall1).
		AddRecallParam(recall2)
	CheckRequestValidate(t, request)

	uri := request.BuildUri()
	fmt.Println(uri.RequestURI())
	got := uri.String()
	want := "be?biz_name=searcher&outfmt=json2&p=testBiz&recall1_trigger_list=1%2C2%2C3&recall2_trigger_list=1%2C1%3B2%2C2%3B3%2C3&return_count=10&s=testBiz"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
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
