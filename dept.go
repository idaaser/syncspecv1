// Package syncspecv1 定义了通用的数据结构, 包括"部门"和"用户"等
package syncspecv1

// Department 定义"部门"的数据结构
type Department struct {
	_ struct{}

	// 唯一标识: 必须, 唯一, 长度<=64
	ID string `json:"id"`

	// 父部门唯一标识: 必须, 长度<=64.
	// 注: Parent为空代表是根部门
	Parent string `json:"parent"`

	// 部门名称: 必须, 非空, 长度<=128
	Name string `json:"name"`

	// 部门在其同级部门的展示顺序, 可选
	Order int `json:"order"`
}

type (

	// PagingDepartments 分页查询返回的部门
	PagingDepartments = PagingResult[*Department]

	// SearchDepartmentRequest 部门搜索请求
	SearchDepartmentRequest struct {
		Keyword string `query:"keyword"`
	}

	// SearchDepartmentResponse 部门搜索响应
	SearchDepartmentResponse struct {
		Data        []*Department `json:"data"`
		ErrResponse `json:",inline"`
	}

	// ListDepatmentRequest 拉取部门列表请求
	ListDepatmentRequest = PagingParam

	// ListDepartmentResponse 拉取部门列表响应
	ListDepartmentResponse struct {
		PagingDepartments `json:",inline"`
		ErrResponse       `json:",inline"`
	}
)
