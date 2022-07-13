package scopes

import (
	"github.com/Geekinn/go-micro/app/forms"
	"gorm.io/gorm"
)

func PageLimit(paginationQuery forms.PaginationQuery) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    pageSize := paginationQuery.GetPageSize()
    return db.Limit(pageSize)
  }
}

func Paginate(paginationQuery forms.PaginationQuery) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    offset := paginationQuery.GetOffset()
    return db.Offset(offset).Scopes(PageLimit(paginationQuery))
  }
}

