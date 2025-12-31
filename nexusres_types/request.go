package nexusres_types

type GinIDRequest struct {
	ID int64 `form:"id" binding:"required"`
}
