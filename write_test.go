package be

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWriteRequest_BuildUri(t *testing.T) {
	writeRequest := NewWriteRequest(WriteTypeAdd, "testTable", "id", map[string]string{})
	writeRequest.AddContent("id", "1234")
	uri := writeRequest.BuildUri()
	fmt.Println(uri.RequestURI())
	got := uri.String()
	want := "sendmsg?content=CMD%3Dadd%1F%0Aid%3D1234%1F%0A&h=2282126479029740061&table=testTable"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v\n want %v", got, want)
	}
}
