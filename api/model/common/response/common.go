package response

type PageResult struct {
	List     interface{} `json:"list"`     // 数据列表
	Total    int64       `json:"total"`    // 总数
	Page     int         `json:"page"`     // 页码
	PageSize int         `json:"pageSize"` // 每页大小
}

// SelectResult 选择列表响应对应对象
type SelectResult struct {
	ID   uint   `json:"id"`   // 主键ID
	Name string `json:"name"` // 名称
}
