package syncspecv1

type (
	// PagingResult 定义了分页请求返回的数据
	PagingResult[T any] struct {
		ErrResponse `json:",inline"`

		// 是否还有待返回的剩余数据, 当为false时表示数据已经全部返回
		HasNext bool `json:"has_next"`
		// 下一次分页请求的游标, 当数据全部返回时, 为""
		Cursor string `json:"cursor"`

		// 本次请求返回的所有数据
		Data []T `json:"data"`
	}

	// PagingParam 定义了分页请求的参数
	PagingParam struct {
		// 单页请求的条数
		Size int `query:"size"`
		// 分页请求的游标, 初始请求为空, 初始请求使用""
		Cursor string `query:"cursor"`
	}
)

// GetSize 若传入的size<=0, 则使用默认的size=50
func (param PagingParam) GetSize() int {
	if param.Size <= 0 {
		return 50
	}
	return param.Size
}
