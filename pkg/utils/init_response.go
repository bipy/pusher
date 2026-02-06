package utils

// Response represents a standard API response
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Resp any    `json:"resp,omitempty"`
}

// SuccessResponse creates a success response with code 0
func SuccessResponse(resp any) Response {
	return Response{
		Code: 0,
		Msg:  "OK",
		Resp: resp,
	}
}

// FailResponse creates a failure response with code 1
func FailResponse(msg string, resp any) Response {
	return Response{
		Code: 1,
		Msg:  msg,
		Resp: resp,
	}
}

// DIYResponse creates a custom response with specified code and message
func DIYResponse(code int, msg string, resp any) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Resp: resp,
	}
}
