package api

import "net/http"

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func Ok(data interface{}) *Response {
	return &Response{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: data,
	}
}
