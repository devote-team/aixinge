package request

import "aixinge/utils/snowflake"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID snowflake.ID `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []snowflake.ID `json:"ids" form:"ids"` //ID数组
}

type IdsRemarkReq struct {
	Ids    []snowflake.ID `json:"ids" form:"ids"`       //ID数组
	Remark string         `json:"remark" form:"remark"` //备注
}

// IdRelIdsReq 一对多关联
type IdRelIdsReq struct {
	RelIds []snowflake.ID `json:"relIds" form:"relIds"` //关联ID数组
	ID     snowflake.ID   `json:"id" form:"id"`         // 主键ID
}

type Empty struct{}
