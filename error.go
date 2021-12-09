package be

import "encoding/json"

type InvalidParamsError struct {
	Message string
}

func (e InvalidParamsError) String() string {
	return e.Message
}

func (e InvalidParamsError) Error() string {
	return e.String()
}

type ClientError struct {
}

type BadResponseError struct {
	RespBody   string
	RespHeader map[string][]string
	HTTPCode   int
}

func (e BadResponseError) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e BadResponseError) Error() string {
	return e.String()
}

func NewBadResponseError(body string, header map[string][]string, httpCode int) *BadResponseError {
	return &BadResponseError{
		RespBody:   body,
		RespHeader: header,
		HTTPCode:   httpCode,
	}
}
