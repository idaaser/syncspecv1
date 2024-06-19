package syncspecv1

// ErrResponse 定义了通用的接口错误返回
type ErrResponse struct {
	// 错误码
	Code string `json:"error"`
	// 错误信息描述
	Msg string `json:"error_message"`
}
