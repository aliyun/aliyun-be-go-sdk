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

func TestClient_ReadVectorFilterClause(t *testing.T) {
	inteTestRead(t, "testdata/test_requests/vector_request.json")
	inteTestRead(t, "testdata/test_requests/vector_request_with_filter.json")
	inteTestRead(t, "testdata/test_requests/x2i_request.json")
	inteTestRead(t, "testdata/test_requests/x2i_request_with_exposure.json")
	inteTestRead(t, "testdata/test_requests/multi_request.json")
}

func TestClient_Write(t *testing.T) {
	tableName := "be-cn-7e22hft5l001_aime_example_expose_2"
	request := NewWriteRequest(WriteTypeAdd, tableName, "user_id", map[string]string{})
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
	tableName := "be-cn-7e22hft5l001_aime_example_expose_2"
	request := NewWriteRequest(WriteTypeDelete, tableName, "user_id", map[string]string{})
	request.AddContent("user_id", "u0003")
	request.AddContent("item_id", "10001")
	request.AddQueryParam("host", "10.0.133.234:80")
	resp, err := client.Write(*request)
	assert.Nil(t, err)

	PrintResult(resp)
	assert.Empty(t, resp)
}

func inteTestRead(t *testing.T, paramPath string) {
	content, rErr := ioutil.ReadFile(paramPath)
	assert.Nil(t, rErr)

	params := &TestReadParams{}
	jErr := json.Unmarshal(content, params)
	assert.Nil(t, jErr)

	readRequest := params.Request
	readRequest.QueryParams = queryParams

	resp, err := client.Read(*readRequest)
	assert.Nil(t, err)
	PrintResult(resp)

	checkers := params.Checkers
	if len(checkers) == 0 {
		return
	}
	for _, checker := range checkers {
		result := checker.check(resp.Result.MatchItems)
		assert.True(t, result, "Failed to check items for checker:"+ToJson(checker))
	}
}

type TestReadParams struct {
	Request  *ReadRequest  `json:"request"`
	Checkers []TestChecker `json:"checkers"`
}

type TestChecker struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

func (c TestChecker) check(items *MatchItem) bool {
	for i := 0; i < items.getResultCount(); i++ {
		item := items.getItems(i)
		itemValue := item[c.Field]
		checkerResult := false
		for _, checkerValue := range c.Values {
			if checkerValue == itemValue {
				checkerResult = true
				break
			}
		}
		if !checkerResult {
			return false
		}
	}
	return true
}
