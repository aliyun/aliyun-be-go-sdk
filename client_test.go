package be

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_Read_Success(t *testing.T) {
	fileName := "testdata/mock_result/read_result_success.json"
	ts, _, err := initHttpTestServer(fileName)
	assert.Nil(t, err, "Failed to init server for:"+fileName)
	testClient := NewClient(ts.URL, "test", "test")

	resp, err := testClient.Read(*initReadRequest())
	assert.Nil(t, err, "Failed to read")
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Result)
	assert.NotNil(t, resp.Result.MatchItems)
	PrintResult(resp)
}

func TestClient_Read_BizNotExist(t *testing.T) {
	fileName := "testdata/mock_result/read_result_biz_not_exist.json"
	ts, tResp, err := initHttpTestServer(fileName)
	assert.Nil(t, err, "Failed to init server for:"+fileName)
	testClient := NewClient(ts.URL, "test", "test")

	_, err = testClient.Read(*initReadRequest())
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "biz not exist")
	respError := err.(*BadResponseError)
	assert.Equal(t, respError.HTTPCode, tResp.Code)
	PrintResult(err)
}

func TestClient_Write_Success(t *testing.T) {
	fileName := "testdata/mock_result/write_result_success.json"
	ts, _, err := initHttpTestServer(fileName)
	assert.Nil(t, err, "Failed to init server for:"+fileName)
	testClient := NewClient(ts.URL, "test", "test")

	resp, err := testClient.Write(*initWriteRequest())
	assert.Nil(t, err)
	PrintResult(resp)
}

func TestClient_Write_Failed(t *testing.T) {
	fileName := "testdata/mock_result/write_result_failed.json"
	ts, tResp, err := initHttpTestServer(fileName)
	assert.Nil(t, err, "Failed to init server for:"+fileName)
	testClient := NewClient(ts.URL, "test", "test")

	_, err = testClient.Write(*initWriteRequest())
	assert.NotNil(t, err)
	respErr := err.(*BadResponseError)
	assert.Equal(t, tResp.Code, respErr.HTTPCode)
	PrintResult(err)
}

func TestClient_WithRequestTimeout(t *testing.T) {
	testClient := NewClient("http://127.0.0.1", "test", "test")
	timeout := 10 * time.Second
	testClient.WithRequestTimeout(timeout)

	assert.Equal(t, timeout, testClient.httpClient.Timeout)
	fmt.Println(testClient.httpClient.Timeout)
}

func TestClient_WithConnectionSize(t *testing.T) {
	testClient := NewClient("http://127.0.0.1", "test", "test")
	timeout := 10 * time.Second
	connectionSize := 333

	testClient.WithRequestTimeout(timeout)
	testClient.WithConnectionSize(connectionSize)

	transportPointer, ok := testClient.httpClient.Transport.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("httpClient.Transport not an *http.Transport"))
	}
	fmt.Printf("timeout: %v, connectionSize: %v", testClient.httpClient.Timeout, transportPointer.MaxIdleConns)
	assert.Equal(t, timeout, testClient.httpClient.Timeout)
	assert.Equal(t, connectionSize, transportPointer.MaxIdleConns)
}

func initWriteRequest() *WriteRequest {
	var content = map[string]string{}
	content["id"] = "10000"
	var contents = []map[string]string{content}
	request := NewWriteRequest(WriteTypeAdd, "tableTable", "id", contents)
	return request
}

func initReadRequest() *ReadRequest {
	request := NewReadRequest("test", 10)
	testRecallParams := NewRecallParam().
		SetTriggerItems([]string{"1:1", "2:1"}).
		SetRecallType(RecallTypeX2I)
	request.AddRecallParam(testRecallParams)
	return request
}

func initHttpTestServer(resultPath string) (*httptest.Server, *HttpTestResponse, error) {
	content, rErr := ioutil.ReadFile(resultPath)
	if rErr != nil {
		return nil, nil, rErr
	}

	testResponse := &HttpTestResponse{}
	jErr := json.Unmarshal(content, testResponse)
	if jErr != nil {
		return nil, nil, jErr
	}

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testResponse.Code)
				w.Write([]byte(testResponse.Body))
			}),
	)
	return ts, testResponse, nil
}

type HttpTestResponse struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}
