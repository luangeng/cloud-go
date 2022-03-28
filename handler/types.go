package handler

const (
	OK    = 0
	ERROR = 1
)

const (
	TRY  = 0
	Good = 1
	End  = 2
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListData struct {
	TotalCount int `json:"totalCount"`
}

func Ok(data interface{}) Response {
	return Response{Code: OK, Message: "ok", Data: data}
}

func Error(msg string) Response {
	return Response{Code: ERROR, Message: msg, Data: nil}
}
