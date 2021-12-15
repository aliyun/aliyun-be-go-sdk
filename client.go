package be

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Endpoint   string
	UserName   string
	PassWord   string
	httpClient *http.Client
}

func NewClient(endpoint string, userName string, passWord string) *Client {
	return &Client{
		Endpoint:   endpoint,
		UserName:   userName,
		PassWord:   passWord,
		httpClient: defaultHttpClient,
	}
}

func (c *Client) Read(readRequest ReadRequest) (*Response, error) {
	vErr := readRequest.Validate()
	if vErr != nil {
		return nil, vErr
	}

	buildUri := readRequest.BuildUri()
	uri := buildUri.RequestURI()
	headers := map[string]string{}

	httpResp, err := request(c, "GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}
	buf, ioErr := ioutil.ReadAll(httpResp.Body)
	if ioErr != nil {
		return nil, NewBadResponseError(ioErr.Error(), httpResp.Header, httpResp.StatusCode)
	}

	readResult := ReadResult{}
	if jErr := json.Unmarshal(buf, &readResult); jErr != nil {
		fmt.Println(jErr)
		return nil, NewBadResponseError("Illegal readResult:"+string(buf), httpResp.Header, httpResp.StatusCode)
	}

	var resp *Response
	if readResult.ErrorCode == 0 {
		result := Result{MatchItems: &readResult.MatchItems}
		resp = NewResponse(result)
	} else {
		return nil, NewBadResponseError(fmt.Sprintf("Failed to read, errorCode[%v], message:%v",
			readResult.ErrorCode, readResult.ErrorMessage), httpResp.Header, httpResp.StatusCode)
	}
	return resp, nil
}

func (c *Client) Write(writeRequest WriteRequest) (*Response, error) {
	vErr := writeRequest.Validate()
	if vErr != nil {
		return nil, vErr
	}
	buildUri := writeRequest.BuildUri()
	uri := buildUri.RequestURI()
	headers := map[string]string{}

	httpResp, err := request(c, "GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	buf, ioErr := ioutil.ReadAll(httpResp.Body)
	if ioErr != nil {
		return nil, NewBadResponseError(ioErr.Error(), httpResp.Header, httpResp.StatusCode)
	}
	writeResult := WriteResult{}
	if jErr := json.Unmarshal(buf, &writeResult); jErr != nil {
		fmt.Println(jErr)
		return nil, NewBadResponseError("Illegal writeResult:"+string(buf), httpResp.Header, httpResp.StatusCode)
	}

	switch writeResult.Errno {
	case 0:
		return NewResponse(Result{}), nil
	case 1:
		return nil, NewBadResponseError(fmt.Sprintf("Failed to write, illegal reqeust body, errorCode[%v], resp:[%v]",
			writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
	default:
		return nil, NewBadResponseError(fmt.Sprintf("Failed to write, errorCode[%v], resp:[%v]",
			writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
	}

}

// Error defines be error
type Error struct {
	HTTPCode int32  `json:"httpCode"`
	Code     string `json:"errorCode"`
	Message  string `json:"errorMessage"`
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
