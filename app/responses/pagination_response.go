package responses

import (
	"fmt"
	"math"

	"github.com/Geekinn/go-micro/app/forms"
	"gorm.io/gorm"
)

type PaginatedResponse struct {
	Total		int64
	PerPage		int
	CurrentPage		int
	LastPage		int
	From		int
	To		int
	Data        any
}


func (PaginatedResponse) Create(data any, paginationQuery forms.PaginationQuery, db *gorm.DB) (PaginatedResponse) {

	var pr PaginatedResponse
	
	var total int64
	if res := db.Count(&total);res.Error != nil{
		fmt.Print(res.Error.Error())
	}
	
	pr.Total = total
	pr.CurrentPage, pr.PerPage = paginationQuery.GetPageAndSize()
	pr.LastPage = int(math.Ceil(float64(total)/float64(pr.PerPage)))
	pr.From = paginationQuery.GetOffset()
	pr.To = pr.From+pr.PerPage
	pr.Data = data

	return pr
}