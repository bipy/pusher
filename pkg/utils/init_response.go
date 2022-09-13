package utils

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Resp any    `json:"resp,omitempty"`
}

func SuccessResponse(resp any) Response {
	return Response{
		Code: 0,
		Msg:  "OK",
		Resp: resp,
	}
}

func FailResponse(msg string, resp any) Response {
	return Response{
		Code: 1,
		Msg:  msg,
		Resp: resp,
	}
}

func DIYResponse(code int, msg string, resp any) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Resp: resp,
	}
}
