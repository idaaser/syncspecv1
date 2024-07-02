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
	// 类似https://datatracker.ietf.org/doc/html/rfc6749#section-4.4.2
	// 不同于oauth2的点在于, Content-Type 不是 application/x-www-form-urlencoded,
	// 改为application/json
	GetTokenRequest struct {
		GrantType    string `json:"grant_type"` // 固定为client_credentials
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}

	// GetTokenResponse 获取token响应
	GetTokenResponse = struct {
		// 接口请求正常时, 返回Tokn
		*Token `json:",inline"`

		// 接口请求失败时, 返回
		ErrResponse `json:",inline"`
	}
)

// Validate 校验请求合法性
func (req GetTokenRequest) Validate() error {
	if req.GrantType != "client_credentials" {
		return fmt.Errorf("grant_type MUST be client_credentials, got: %q", req.GrantType)
	}

	if req.ClientID == "" {
		return fmt.Errorf("client id MUST NOT be empty")
	}

	if req.ClientSecret == "" {
		return fmt.Errorf("client secret MUST NOT be empty")
	}

	return nil
}
