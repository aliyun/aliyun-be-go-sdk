package be

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var client = NewClient("http://curl-proxy-vpc-pre.airec.aliyun-inc.com", "test", "test")

func TestClient_Read(t *testing.T) {
	bizName := "silan_test_pai"
	returnCount := 10
	request := NewReadRequest(bizName, returnCount)
	testRecallParams := NewRecallParam().
		SetTriggerItems([]string{"1:1", "2:1"}).
		SetRecallType(RecallTypeX2i)
	request.AddRecallParam(testRecallParams).
		AddQueryParam("host", "10.2.207.13:10579")

	resp, err := client.Read(*request)
	if err != nil {
		t.Errorf("Failed to read, %s", err)
	}
	respByte, _ := json.Marshal(resp)
	fmt.Println(string(respByte))
	if reflect.DeepEqual(resp.Result, Result{}) {
		t.Errorf("Empty resp")
	}
	if resp.Result.MatchItems == nil {
		t.Errorf("Empty match items")
	}
}

func TestClient_Write(t *testing.T) {
	tableName := "silantest"
	request := NewWriteRequest(WriteTypeAdd, tableName, "id", map[string]string{})
	request.AddContent("id", "10000")
	request.AddQueryParam("host", "10.0.133.234:80")
	_, err := client.Write(*request)
	if err != nil {
		t.Errorf("Failed to write, %s", err)
	}
}
