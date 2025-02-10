package syncspecv1

import "fmt"

// Group 定义"group"的数据结构
type Group struct {
	_ struct{}

	// ID 唯一标识: 必须, 唯一, 长度<=64
	ID string `json:"id"`

	// Name group名称: 必须, 唯一, 非空, 长度<=128
	Name string `json:"name"`
}

type (

	// PagingGroups 分页查询返回的group
	PagingGroups = PagingResult[*Group]

	// SearchGroupRequest group搜索请求
	SearchGroupRequest struct {
		Keyword string `query:"keyword"`
	}

	// SearchGroupResponse group搜索响应
	SearchGroupResponse struct {
		Data        []*Group `json:"data"`
		ErrResponse `json:",inline"`
	}

	// ListGroupRequest 拉取group列表请求
	ListGroupRequest = PagingParam

	// ListGroupResponse 拉取group列表响应
	ListGroupResponse struct {
		PagingGroups `json:",inline"`
		ErrResponse  `json:",inline"`
	}
	// ListGroupMembershipRequest 拉取指定group下的用户列表请求
	ListGroupMembershipRequest struct {
		Group string `query:"id"`
		PagingParam
	}

	// ListGroupMembershipReseponse 拉取指定group下的用户列表响应, 仅需要返回用户id
	ListGroupMembershipResponse struct {
		Members     PagingResult[string] `json:",inline"`
		ErrResponse `json:",inline"`
	}
)

func (d Group) String() string {
	return fmt.Sprintf("id=%q, name=%q", d.ID, d.Name)
}
