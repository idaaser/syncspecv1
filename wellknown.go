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

	// 部门列表接口地址: 通过分页的方式返回部门列表
	ListDepartmentsEndpoint string `json:"list_department_endpoint"`
	// 部门搜索接口地址: 根据关键字搜索部门
	SearchDepartmentEndpoint string `json:"search_department_endpoint"`

	// 获取指定部门下的直属用户详情接口: 通过分页的方式返回部门下直属用户列表
	ListUsersInDeptEndpoint string `json:"list_deptartment_users_endpoint"`
	// 用户搜索接口地址: 根据关键字搜索用户
	SearchUserEndpoint string `json:"search_user_endpoint"`

	// group列表接口地址: 通过分页的方式返回group列表
	ListGroupsEndpoint string `json:"list_group_endpoint"`
	// 返回指定group下的用户id
	ListUsersInGroupEndpoint string `json:"list_group_users_endpoint"`
	// group搜索接口地址: 根据关键字搜索group
	SearchGroupEndpoint string `json:"search_group_endpoint"`
}

// Spec 返回支持的协议版本号
func (w Wellknown) Spec() string {
	return "v1"
}

// MarshalJSON 自定义json序列化, 添加spec字段为v1
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
	if err := w.validateURL("token_endpoint", w.TokenEndpoint, true); err != nil {
		return err
	}

	if err := w.validateURL("list_deptartment_users_endpoint",
		w.ListUsersInDeptEndpoint, true); err != nil {
		return err
	}

	if err := w.validateURL("search_user_endpoint",
		w.SearchUserEndpoint, true); err != nil {
		return err
	}

	if err := w.validateURL("list_department_endpoint",
		w.ListDepartmentsEndpoint, true); err != nil {
		return err
	}

	if err := w.validateURL("search_department_endpoint",
		w.SearchDepartmentEndpoint, true); err != nil {
		return err
	}

	if err := w.validateURL("list_group_endpoint",
		w.ListGroupsEndpoint, false); err != nil {
		return err
	}

	if err := w.validateURL("list_group_users_endpoint",
		w.ListUsersInGroupEndpoint, false); err != nil {
		return err
	}

	if err := w.validateURL("search_group_endpoint",
		w.SearchGroupEndpoint, false); err != nil {
		return err
	}

	return nil
}

func (w Wellknown) validateURL(field, u string, required bool) error {
	if required && u == "" {
		return fmt.Errorf("%s MUST NOT be empty", field)
	}

	if _, err := url.Parse(u); err != nil {
		return err
	}

	return nil
}
