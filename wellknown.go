package syncspec

import "encoding/json"

// Wellknown 定义了服务端的公开信息
type Wellknown struct {
	TokenEndpoint string `json:"token_endpoint"`
	// 支持的token接口鉴权方式"client_secret_basic", "client_secret_post"
	TokenEndpointAuthMethods []string `json:"token_endpoint_auth_methods_supported"`

	// 获取指定部门下的直属用户接口: 通过分页的方式返回部门下直属用户列表
	ListUsersInDeptEndpoint string `json:"list_deptartment_users_endpoint"`
	// 用户搜索接口地址: 根据关键字搜索用户
	SearchUserEndpoint string `json:"search_user_endpoint"`

	// 部门列表接口地址: 通过分页的方式返回部门列表
	ListDepartmentsEndpoint string `json:"list_department_endpoint"`
	// 部门搜索接口地址: 根据关键字搜索部门
	SearchDepartmentEndpoint string `json:"search_department_endpoint"`
}

// Spec 返回支持的协议版本号
func (w Wellknown) Spec() string {
	return "v1"
}

const (
	// TokenEndpointAuthMethodBasic basic auth的鉴权方式
	TokenEndpointAuthMethodBasic = "client_secret_basic"
	// TokenEndpointAuthMethodPost 表单提交的方式
	TokenEndpointAuthMethodPost = "client_secret_post"
)

// MarshalJSON json序列化, 添加spec字段为v1
func (w Wellknown) MarshalJSON() ([]byte, error) {
	type Alias Wellknown

	return json.Marshal(
		&struct {
			Spec string `json:"spec"` // 支持的协议版本号, 固定为v1
			Alias
		}{
			Spec:  w.Spec(),
			Alias: (Alias)(w),
		})
}
