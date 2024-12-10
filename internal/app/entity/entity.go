package entity

type PaginatedResult struct {
	Data       interface{} `json:"data"`        // 查询结果
	TotalCount int64       `json:"total_count"` // 总记录数
	Page       int         `json:"page"`        // 当前页码
	PageSize   int         `json:"page_size"`   // 每页大小
	TotalPage  int         `json:"total_size"`
	NextPage   int         `json:"next_page"`
}
