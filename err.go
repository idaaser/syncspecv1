package syncspecv1

import "fmt"

// ErrResponse 定义了通用的接口错误返回
type ErrResponse struct {
	// 错误码, 当为""时表示请求正常
	Code string `json:"code,omitempty"`
	// 错误信息描述
	Msg string `json:"msg,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// IsError 判断是否是请求错误
func (err *ErrResponse) IsError() bool {
	return err.Code != ""
}

func (err *ErrResponse) Error() string {
	return fmt.Sprintf("code=%s, msg=%s, request_id=%s",
		err.Code, err.Msg, err.RequestID,
	)
}

// 定义常见的error Code
const (
	// 非法请求, 通常伴随http status code 400一起使用
	ErrInvalidRequest = "invalid_request"

	// 接口调用时使用的acces_token不合法, 通常伴随http status code 401一起使用
	ErrInvalidToken = "invalid_token"

	// 请求access_token时, 使用的client_id/client_secret不正确, 通常伴随http status code 401一起使用
	ErrInvalidClient = "invalid_client"
)
