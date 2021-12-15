package be

import (
	"fmt"
	"hash/fnv"
	"net/url"
	"strconv"
	"strings"
)

type WriteType string

const (
	WriteTypeAdd    WriteType = "ADD"
	WriteTypeDelete WriteType = "DELETE"
)

type WriteRequest struct {
	WriteType    WriteType         `json:"write_type"`
	TableName    string            `json:"table_name"`
	Contents     map[string]string `json:"contents"`
	PartitionKey string            `json:"partition_key"`
	QueryParams  map[string]string `json:"query_params"`
}

func NewWriteRequest(writeType WriteType, tableName string, partitionKey string, contents map[string]string) *WriteRequest {
	return &WriteRequest{WriteType: writeType,
		TableName:    tableName,
		PartitionKey: partitionKey,
		Contents:     contents,
		QueryParams:  map[string]string{},
	}
}

func (r *WriteRequest) AddContent(key string, value string) *WriteRequest {
	r.Contents[key] = value
	return r
}

func (r *WriteRequest) Validate() error {
	if r.TableName == "" {
		return InvalidParamsError{"Empty table name"}
	}
	if len(r.Contents) == 0 {
		return InvalidParamsError{"Empty contents"}
	}
	if r.PartitionKey == "" {
		return InvalidParamsError{"Partition key not set"}
	}
	partitionKeyExist := false
	for k, v := range r.Contents {
		if k == "" || v == "" {
			return InvalidParamsError{fmt.Sprintf("Key or value is empty for kv pair[%s][%s]", k, v)}
		}
		if k == r.PartitionKey {
			partitionKeyExist = true
		}
	}
	if !partitionKeyExist {
		return InvalidParamsError{fmt.Sprintf("Partition key[%s] not exist in contents", r.PartitionKey)}
	}
	return nil
}

func (r *WriteRequest) AddQueryParam(key string, value string) *WriteRequest {
	r.QueryParams[key] = value
	return r
}

func (r *WriteRequest) SetQueryParams(params map[string]string) *WriteRequest {
	r.QueryParams = params
	return r
}

func (r *WriteRequest) BuildUri() url.URL {
	uri := url.URL{Path: "sendmsg"}

	var builder strings.Builder
	var separator byte = 31
	var lineBreak byte = '\n'

	builder.WriteString("CMD=")
	builder.WriteString(strings.ToLower(string(r.WriteType)))
	builder.WriteByte(separator)
	builder.WriteByte(lineBreak)

	var partitionKeyValue string
	for k, v := range r.Contents {
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(v)
		builder.WriteByte(separator)
		builder.WriteByte(lineBreak)
		if r.PartitionKey == k {
			partitionKeyValue = v
		}
	}

	h := fnv.New64a()
	_, err := h.Write([]byte(partitionKeyValue))
	if err != nil {
		//TODO add log
	}
	hashValue := h.Sum64()

	query := uri.Query()
	query.Add("table", r.TableName)
	query.Add("h", strconv.Itoa(int(hashValue)))
	query.Add("msg", builder.String())

	rawQuery := query.Encode()
	for k, v := range r.QueryParams {
		rawQuery = rawQuery + "&" + k + "=" + v
	}
	uri.RawQuery = rawQuery
	return uri
}
