package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

//对错误信息统一标准化
var (
	ErrorRequestBodyParseFailed = ErrorResponse{400, Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{401, Err{Error: "User Auth failed", ErrorCode: "002"}}
	//
)
