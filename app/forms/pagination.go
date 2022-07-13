package forms

//Pagination ...
type PaginationQuery struct {
	Page int `form:"page"`
	PageSize int `form:"pagesize"`
}

func (paginationQuery PaginationQuery)GetPage() (int)  {
  page := paginationQuery.Page
    if page == 0 {
      page = 1
    }
    return page
}

func (paginationQuery PaginationQuery)GetPageSize() (int)  {
    pageSize := paginationQuery.PageSize
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }
    return pageSize
}

func (paginationQuery PaginationQuery)GetPageAndSize() (int, int)  {
  return paginationQuery.GetPage(),paginationQuery.GetPageSize()
}

func (paginationQuery PaginationQuery)GetOffset()(int)  {
  page, pageSize := paginationQuery.GetPageAndSize()
  return (page - 1) * pageSize
}