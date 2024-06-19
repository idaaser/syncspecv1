package syncspecv1

import (
	"encoding/json"
	"fmt"
)

// Token 定义了接口鉴权需要的token
// 遵循oauth2的定义, https://datatracker.ietf.org/doc/html/rfc6749#section-4.4.3
type Token struct {
	AccessToken string `json:"access_token"`

	// 有效期, 单位秒, 比如7200,表示2小时
	ExpiresIn int32 `json:"expires_in"`
}

// TokenType token的类型,固定为Bearer
func (t Token) TokenType() string {
	return "Bearer"
}

// MarshalJSON json序列化, 添加token_type字段为Bearer
func (t Token) MarshalJSON() ([]byte, error) {
	type Alias Token

	return json.Marshal(
		&struct {
			TokenType string `json:"token_type"`
			Alias
		}{
			TokenType: t.TokenType(),
			Alias:     (Alias)(t),
		})
}

type (
	// GetTokenRequest 获取token请求
	GetTokenRequest struct {
		ClientID     string `form:"client_id"`
		ClientSecret string `form:"client_secret"`
	}

	// GetTokenResponse 获取token响应
	GetTokenResponse = Token
)

// Validate 校验请求合法性
func (req GetTokenRequest) Validate() error {
	if req.ClientID == "" {
		return fmt.Errorf("client id MUST NOT be empty")
	}

	if req.ClientSecret == "" {
		return fmt.Errorf("client secret MUST NOT be empty")
	}

	return nil
}
