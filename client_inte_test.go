package be

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var client = NewClient("http://curl-proxy-vpc-pre.airec.aliyun-inc.com", "test", "test")
var queryParams = map[string]string{}

func TestMain(m *testing.M) {
	queryParams["host"] = "10.2.207.14:16832"
	m.Run()
}

func TestClient_Read_Vector(t *testing.T) {
	fileName := "testdata/inte_test_requests/vector_request.json"
	request, err := initTestReadRequest(fileName)
	assert.Nil(t, err, "Failed to init request from file:"+fileName)
	resp, err := client.Read(*request)
	assert.Nil(t, err)
	assert.Equal(t, request.ReturnCount, len(resp.Result.MatchItems.FieldValues))
}

func TestClient_Read_VectorClause(t *testing.T) {
	fileName := "testdata/inte_test_requests/vector_request_with_filter.json"
	request, err := initTestReadRequest(fileName)
	assert.Nil(t, err, "Failed to init request from file:"+fileName)
	resp, err := client.Read(*request)
	assert.Nil(t, err)
	items := resp.Result.MatchItems
	for i := 0; i < items.getResultCount(); i++ {
		assert.Equal(t, "100", items.getItems(i)["field1"])
	}

}

func TestClient_Write(t *testing.T) {
	tableName := "aime_example_expose_2"
	instanceName := "be-cn-7e22hft5l001"
	request := NewWriteRequest(WriteTypeAdd, instanceName, tableName, "user_id", map[string]string{})
	request.AddContent("user_id", "u0003")
	request.AddContent("item_id", "10001")
	request.AddContent("time", "1640242662")
	request.AddQueryParam("host", "10.0.133.234:80")
	resp, err := client.Write(*request)
	assert.Nil(t, err)

	PrintResult(resp)
	assert.Empty(t, resp)
}

func TestClient_Write_Delete(t *testing.T) {
	tableName := "aime_example_expose_2"
	instanceName := "be-cn-7e22hft5l001"
	request := NewWriteRequest(WriteTypeDelete, instanceName, tableName, "user_id", map[string]string{})
	request.AddContent("user_id", "u0003")
	request.AddContent("item_id", "10001")
	request.AddQueryParam("host", "10.0.133.234:80")
	resp, err := client.Write(*request)
	assert.Nil(t, err)

	PrintResult(resp)
	assert.Empty(t, resp)
}

func initTestReadRequest(requestFilePath string) (*ReadRequest, error) {
	content, rErr := ioutil.ReadFile(requestFilePath)
	if rErr != nil {
		return nil, rErr
	}

	readRequest := &ReadRequest{}
	jErr := json.Unmarshal(content, readRequest)
	if jErr != nil {
		return nil, jErr
	}
	readRequest.QueryParams = queryParams
	PrintResult(readRequest)
	return readRequest, nil
}
