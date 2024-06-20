package syncspecv1

// ErrResponse 定义了通用的接口错误返回
type ErrResponse struct {
	// 错误码
	Code string `json:"error,omitempty"`
	// 错误信息描述
	Msg string `json:"error_message,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`
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
