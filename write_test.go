package be

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteRequest_BuildUri(t *testing.T) {
	writeRequest := NewWriteRequest(WriteTypeAdd, "testInstance", "testTable", "id", map[string]string{})
	writeRequest.AddContent("id", "1234")
	uri := writeRequest.BuildUri()
	fmt.Println(uri.RequestURI())
	got := uri.String()
	want := "sendmsg?h=2282126479029740061&msg=CMD%3Dadd%1F%0Aid%3D1234%1F%0A&table=testInstance_testTable"
	assert.Equal(t, got, want)
}
