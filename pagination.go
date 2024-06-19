package syncspecv1

type (
	// PagingResult 定义了分页请求返回的数据
	PagingResult[T any] struct {
		ErrResponse `json:",inline"`

		// 是否还有待返回的剩余数据, 当为false时表示数据已经全部返回
		HasNext bool `json:"has_next"`
		// 下一次分页请求的游标
		Cursor string `json:"cursor"`

		// 本次请求返回的所有数据
		Data []T `json:"data"`
	}

	// PagingParam 定义了分页请求的参数
	PagingParam struct {
		// 单页请求的最大条数
		Size int `json:"size"`
		// 分页请求的游标, 初始请求为空
		Cursor string `json:"cursor"`
	}
)
