package be

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// WithRequestTimeout with custom timeout for a request
func (c *Client) WithRequestTimeout(timeout time.Duration) *Client {
	c.httpClient.Timeout = timeout
	return c
}

func (c *Client) WithConnectionSize(connectionCount int) *Client {
	transportTripper := c.httpClient.Transport
	transportPointer, ok := transportTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("httpClient.Transport not an *http.Transport"))
	}

	transport := *transportPointer
	transport.MaxIdleConnsPerHost = connectionCount
	transport.MaxIdleConns = connectionCount

	c.httpClient = &http.Client{
		Timeout:   c.httpClient.Timeout,
		Transport: &transport,
	}
	return c
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
	defer httpResp.Body.Close()

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

	// TODO modify to batch write
	for i := 0; i < len(writeRequest.Contents); i++ {
		buildUri := writeRequest.BuildUri(i)
		uri := buildUri.RequestURI()
		headers := map[string]string{}

		httpResp, err := request(c, "GET", uri, headers, nil)
		if err != nil {
			return nil, err
		}
		// defer httpResp.Body.Close()

		buf, ioErr := ioutil.ReadAll(httpResp.Body)
		httpResp.Body.Close()
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
			continue
		case 1:
			return nil, NewBadResponseError(fmt.Sprintf("Failed to write, illegal reqeust body, errorCode[%v], resp:[%v]",
				writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
		default:
			return nil, NewBadResponseError(fmt.Sprintf("Failed to write, errorCode[%v], resp:[%v]",
				writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
		}
	}
	return NewResponse(Result{}), nil

}
