package syncspecv1

import (
	"fmt"
)

// User 定义"用户"的数据结构
type User struct {
	_ struct{}

	// 唯一标识: 必须, 唯一, 长度<=64
	ID string `json:"id"`

	// 登录名: 唯一, 长度<=64, Username、Email、Mobile 这三个字段建议得有一个
	Username *string `json:"username"`

	// 邮箱: 唯一, Username、Email、Mobile 这三个字段建议得有一个
	Email *string `json:"email"`

	// 手机号: 唯一, Username、Email、Mobile 这三个字段建议得有一个
	// 注: 格式为E.164格式(https://en.wikipedia.org/wiki/E.164), 比如+8613411112222
	Mobile *string `json:"mobile"`

	// 显示名或姓名: 必须, 长度<=64
	Name string `json:"name"`

	// 职位: 可选, 长度<=64
	Position *string `json:"position"`

	// 员工工号: 可选, 长度<=64
	EmployeeNumber *string `json:"employee_number"`

	// 头像URL: 可选
	Avatar *string `json:"avatar"`

	// 入职时间戳(unix timestamp): 可选
	JoinTime int64 `json:"join_time"`

	// 用户状态: 启用或禁用
	Active bool `json:"active"`

	// 用户所属主部门唯一标识: 必须
	MainDepartmentID string `json:"main_department"`

	// 用户所属其他部门唯一标识(除去MainDepartmentID以外): 可选
	OtherDepartmentsID []string `json:"other_departments"`

	// 用户在其部门下的展示顺序, 可选
	Order int `json:"order"`

	// 其他字段: 可选
	ExtAttrs map[string]any `json:"extattrs"`
}

type (
	// PagingUsers 分页查询返回的用户
	PagingUsers = PagingResult[*User]

	// SearchUserRequest 部门搜索请求
	SearchUserRequest struct {
		Keyword string `query:"keyword"`
	}

	// SearchUserResponse 用户搜索响应
	SearchUserResponse struct {
		Data        []*User `json:"data"`
		ErrResponse `json:",inline"`
	}

	// ListUsersInDepatmentRequest 拉取部门直属用户列表请求
	ListUsersInDepatmentRequest struct {
		DepartmentID string `query:"id"`
		PagingParam
	}
	// ListUsersInDepartmentResponse 拉取部门直属用户列表响应
	ListUsersInDepartmentResponse struct {
		PagingUsers `json:",inline"`
		ErrResponse `json:",inline"`
	}
)

// Validate 合法性检查
func (req ListUsersInDepatmentRequest) Validate() error {
	if req.DepartmentID == "" {
		return fmt.Errorf("query param of deptid MUST NOT be empty")
	}
	return nil
}
