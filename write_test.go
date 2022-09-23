package be

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteRequest_BuildUri(t *testing.T) {
	content := map[string]string{}
	content["id"] = "1234"
	contents := []map[string]string{content}
	writeRequest := NewWriteRequest(WriteTypeAdd, "testTable", "id", contents)
	uri := writeRequest.BuildUri(0)
	fmt.Println(uri.RequestURI())
	got := uri.String()
	want := "sendmsg?h=2282126479029740061&msg=CMD%3Dadd%1F%0Aid%3D1234%1F%0A&table=testTable"
	assert.Equal(t, got, want)
}
