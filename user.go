package syncspecv1

import "time"

// User 定义"用户"的数据结构
type User struct {
	// 唯一标识: 必须, 唯一, 长度<=64
	ID string `json:"id"`

	// 登录名: 唯一, 长度<=64, Username、Email、Mobile 这三个字段建议得有一个
	Username *string `json:"username,omitempty"`

	// 邮箱: 唯一, Username、Email、Mobile 这三个字段建议得有一个
	Email *string `json:"email,omitempty"`

	// 手机号: 唯一, Username、Email、Mobile 这三个字段建议得有一个
	// 注: 格式为E.164格式(https://en.wikipedia.org/wiki/E.164), 比如+8613411112222
	Mobile *string `json:"mobile,omitempty"`

	// 显示名或姓名: 必须, 长度<=64
	Name string `json:"name"`

	// 职位: 可选, 长度<=64
	Position *string `json:"position,omitempty"`

	// 员工工号: 可选, 长度<=64
	EmployeeNumber *string `json:"employee_number,omitempty"`

	// 头像URL: 可选
	Avatar *string `json:"avatar,omitempty"`

	// 入职时间: 可选
	JoinTime time.Time `json:"join_time"`

	Status UserStatus `json:"status"`

	// 用户所属主部门唯一标识: 必须
	MainDepartmentID string `json:"main_department"`

	// 用户所属其他部门唯一标识(除去MainDepartmentID以外): 可选
	OtherDepartmentsID []string `json:"other_departments"`

	// 其他字段: 可选
	ExtAttrs map[string]any `json:"extattrs"`
}

// UserStatus 用户状态位
type UserStatus int32

// 已知的用户状态位
const (
	// 禁用
	UserStatusDisabled UserStatus = 0
	// 待激活
	UserStatusInitialized UserStatus = 1
	// 启用
	UserStatusEnabled UserStatus = 2
)

type (
	// PagingUsers 分页查询返回的用户
	PagingUsers = PagingResult[*User]

	// SearchUserRequest 部门搜索请求
	SearchUserRequest struct {
		Keyword string `json:"keyword"`
	}

	// SearchUserResponse 用户搜索响应
	SearchUserResponse struct {
		Data        []*User `json:"data"`
		ErrResponse `json:",inline"`
	}

	// ListUsersInDepatmentRequest 拉取部门直属用户列表请求
	ListUsersInDepatmentRequest struct {
		DepartmentID string `json:"deptid"`
		PagingParam  `json:",inline"`
	}
	// ListUsersInDepartmentResponse 拉取部门直属用户列表响应
	ListUsersInDepartmentResponse = PagingUsers
)
