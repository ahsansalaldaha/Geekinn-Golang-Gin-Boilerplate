package scopes

import (
	"github.com/Geekinn/go-micro/app/forms"
	"gorm.io/gorm"
)


func Paginate(paginationQuery forms.PaginationQuery) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    
    page := paginationQuery.Page
    if page == 0 {
      page = 1
    }

    pageSize := paginationQuery.PageSize
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

