package forms

//Pagination ...
type PaginationQuery struct {
	Page int `form:"page"`
	PageSize int `form:"pagesize"`
}