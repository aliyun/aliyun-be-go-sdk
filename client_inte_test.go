package be

import (
	"encoding/json"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

//var client = NewClient("http://curl-proxy-vpc-pre.airec.aliyun-inc.com", "test", "test")
var client = NewClient("http://curl-proxy-cn-hangzhou.vpc.airec.aliyun-inc.com", "test", "test")
var queryParams = map[string]string{}

func TestMain(m *testing.M) {
	queryParams["host"] = "10.0.187.27:18032"
	m.Run()
}

func TestShihuo_Read(t *testing.T) {
	inteTestRead(t, "testdata/test_requests/shihuo_request.json")
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
	var content1 = map[string]string{}
	content1["user_id"] = "u0003"
	content1["item_id"] = "10001"
	content1["time"] = "1640242663"

	var content2 = map[string]string{}
	content2["user_id"] = "u0003"
	content2["item_id"] = "10002"
	content2["time"] = "1640242664"

	var contents = []map[string]string{content1, content2}
	request := NewWriteRequest(WriteTypeAdd, tableName, "user_id", contents)
	request.AddQueryParam("host", "10.0.133.234:80")

	resp, err := client.Write(*request)
	assert.Nil(t, err)

	PrintResult(resp)
	assert.Empty(t, resp)
}

func TestClient_Write_Delete(t *testing.T) {
	tableName := "be-cn-7e22hft5l001_aime_example_expose_2"
	var content = map[string]string{}
	content["user_id"] = "u0003"
	content["item_id"] = "10001"
	content["host"] = "10.0.133.234:80"
	var contents = []map[string]string{content}

	request := NewWriteRequest(WriteTypeDelete, tableName, "user_id", contents)

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
	for k, v := range queryParams {
		readRequest.QueryParams[k] = v
	}

	resp, err := client.Read(*readRequest)
	assert.Nil(t, err)
	PrintResult(resp)
	item := resp.Result.MatchItems.getItems(0)
	itemId := item["item_id"]
	intValue, _ := itemId.(json.Number).Int64()
	fmt.Printf("type:%s value:%d\n", reflect.TypeOf(intValue), intValue)

	tagsRules := item["tag_rules"]
	fmt.Printf("type:%s， value:%s\n", reflect.TypeOf(tagsRules), tagsRules)

	tags := item["new_style_tags"]
	fmt.Printf("type:%s\n", reflect.TypeOf(tags))

	//fmt.Printf("type:%s\n", reflect.TypeOf(.Int64()))

	checkers := params.Checkers
	if len(checkers) == 0 {
		return
	}
	for _, checker := range checkers {
		result := checker.check(resp.Result.MatchItems)
		assert.True(t, result, "Failed to check items for checker:"+ToJson(checker))
	}
}

func TestReadBenchmark(b *testing.T) {
	client = NewClient("http://shihuo-pre.public.be.aliyuncs.com", "aliyun-rec", "Rec1234#")
	client.EnableMetric = true
	client.InitMetrics()
	timer := metrics.NewTimer()
	metrics.GetOrRegister("timer.read", timer)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())

	params := readParams("testdata/test_requests/param.txt")
	request := NewReadRequest("shihuo_common", 2000)
	request.IsRawRequest = true
	request.SetQueryParams(params)
	request.OutFmt = "fb2"

	for i := 0; i < 1000; i++ {
		start := time.Now()
		_, err := client.Read(*request)
		check(err)
		timer.Update(time.Since(start))
	}
}

func TestGetByPb(t *testing.T) {
	client = NewClient("http://shihuo-pre.public.be.aliyuncs.com", "aliyun-rec", "Rec1234#")
	params := readParams("testdata/test_requests/param.txt")
	//params := readParams("testdata/test_requests/empty_result_param.txt")
	request := NewReadRequest("shihuo_common", 20)
	request.IsRawRequest = true
	request.SetQueryParams(params)
	request.OutFmt = "fb2"
	request.IsPost = true

	resp, err := client.Read(*request)
	check(err)
	PrintResult(resp)
}

func TestGetByPbWithTrace(t *testing.T) {
	client = NewClient("http://shihuo-pre.public.be.aliyuncs.com", "aliyun-rec", "Rec1234#")
	params := readParams("testdata/test_requests/param_with_trace.txt")
	request := NewReadRequest("shihuo_common", 1)
	request.IsRawRequest = true
	request.SetQueryParams(params)
	request.OutFmt = "fb2"
	request.IsPost = true

	resp, err := client.Read(*request)
	check(err)
	PrintResult(resp)
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

func readParams(filePath string) map[string]string {
	content, err := os.ReadFile(filePath)
	check(err)
	paramsString := string(content)
	kvPairs := strings.Split(paramsString, "&")
	params := make(map[string]string)
	for _, kvPair := range kvPairs {
		kv := strings.Split(kvPair, "=")
		params[kv[0]] = kv[1]
	}
	return params
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
