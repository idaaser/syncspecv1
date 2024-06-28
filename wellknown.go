package syncspecv1

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Wellknown 定义了服务端的公开信息
type Wellknown struct {
	// 获取鉴权access_token的接口
	TokenEndpoint string `json:"token_endpoint"`

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

// Validate 校验合法性
func (w Wellknown) Validate() error {
	if err := w.validateURL("token_endpoint", w.TokenEndpoint); err != nil {
		return err
	}

	if err := w.validateURL("list_deptartment_users_endpoint",
		w.ListUsersInDeptEndpoint); err != nil {
		return err
	}

	if err := w.validateURL("search_user_endpoint",
		w.SearchUserEndpoint); err != nil {
		return err
	}

	if err := w.validateURL("list_department_endpoint",
		w.ListDepartmentsEndpoint); err != nil {
		return err
	}

	if err := w.validateURL("search_department_endpoint",
		w.SearchDepartmentEndpoint); err != nil {
		return err
	}

	return nil
}

func (w Wellknown) validateURL(field, u string) error {
	if u == "" {
		return fmt.Errorf("%s MUST NOT be empty", field)
	}

	if _, err := url.Parse(u); err != nil {
		return err
	}

	return nil
}
