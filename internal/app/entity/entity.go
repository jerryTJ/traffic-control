package entity

type PaginatedResult struct {
	Data       interface{} `json:"data"`        // 查询结果
	TotalCount int64       `json:"total_count"` // 总记录数
	Page       int         `json:"page"`        // 当前页码
	PageSize   int         `json:"page_size"`   // 每页大小
	TotalPage  int         `json:"total_size"`  // 总页数
	NextPage   int         `json:"next_page"`   //下一页
}

type ChainVo struct {
	ID         uint           `json:"id"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  string         `json:"updated_at"`
	Name       string         `json:"name"`
	Version    string         `json:"version"`
	ServerInfo []ServerInfoVo `json:"server_infos"`
}

type ServerInfoVo struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	Rank uint   `json:"rank"`
}
type ResponseVo struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
