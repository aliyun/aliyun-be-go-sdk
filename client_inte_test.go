package be

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err)

	respByte, _ := json.Marshal(resp)
	fmt.Println(string(respByte))
	assert.NotEmpty(t, resp.Result)
	assert.NotNil(t, resp.Result.MatchItems)
}

func TestClient_Write(t *testing.T) {
	tableName := "silantest"
	request := NewWriteRequest(WriteTypeAdd, tableName, "id", map[string]string{})
	request.AddContent("id", "10000")
	request.AddQueryParam("host", "10.0.133.234:80")
	resp, err := client.Write(*request)
	assert.Nil(t, err)

	PrintResult(resp)
	assert.Empty(t, resp)
}
