package be

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Endpoint   string
	UserName   string
	PassWord   string
	httpClient *http.Client
}

func (c *Client) read(readRequest ReadRequest) (*Response, error) {
	buildUri := readRequest.BuildUri()
	uri := buildUri.RequestURI()
	headers := map[string]string{}
	_, err := request(c, "GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) write(request WriteRequest) (*Response, error) {
	return nil, nil
}

// Error defines be error
type Error struct {
	HTTPCode  int32  `json:"httpCode"`
	Code      string `json:"errorCode"`
	Message   string `json:"errorMessage"`
	RequestID string `json:"requestID"`
}

// NewClientError new client error
func NewClientError(err error) *Error {
	if err == nil {
		return nil
	}
	if clientError, ok := err.(*Error); ok {
		return clientError
	}
	clientError := new(Error)
	clientError.HTTPCode = -1
	clientError.Code = "ClientError"
	clientError.Message = err.Error()
	return clientError
}

func (e Error) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e Error) Error() string {
	return e.String()
}
