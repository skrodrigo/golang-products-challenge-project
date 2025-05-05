package shared

type PageRequest struct {
	Limit  int    `json:"limit" form:"limit" binding:"required"`
	OffSet int    `json:"offset" form:"offset" binding:"required"`
	Sort   string `json:"sort" form:"sort" binding:"required"`
	Desc   bool   `json:"desc" form:"desc" binding:"required"`
}
