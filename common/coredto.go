package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseOfErr(err error) *Response {
	return &Response{
		Code:    -1,
		Message: err.Error(),
		Data:    nil,
	}
}

func NewResponseOfSuccess(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "",
		Data:    data,
	}
}
