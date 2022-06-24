package models

import (
	"github.com/Geekinn/go-micro/database"
	"github.com/Geekinn/go-micro/app/forms"
	"gorm.io/gorm"
)

//Article ...
type Article struct {
	gorm.Model
	ID        int64    `gorm:"primaryKey"`
	UserID    int64
	Title     string
	Content   string
	UpdatedAt int64    `gorm:"autoUpdateTime"`
	CreatedAt int64	   `gorm:"autoCreateTime"`
}

//ArticleModel ...
type ArticleModel struct{}

func (m ArticleModel) Migrate(){
	db.GetDB().AutoMigrate(&Article{})
}

//Create ...
func (m ArticleModel) Create(userID int64, form forms.CreateArticleForm) (articleID int64, err error) {

	article := Article{Title: form.Title, Content: form.Content, UserID: userID}
	if dbc := db.GetDB().Create(&article); dbc.Error != nil {
		return articleID, dbc.Error
	}else{
		return article.ID, dbc.Error
	}
}

//One ...
func (m ArticleModel) One(userID, id int64) (article Article, err error) {

	if dbc := db.GetDB().Where("user_id = ?", userID).Where(id).First(&article); dbc.Error != nil {
		return article, dbc.Error
	}else{
		return article, dbc.Error
	}
}

//All ...
func (m ArticleModel) All(userID int64) (articles []Article, err error) {
	if dbc := db.GetDB().Where("user_id = ?", userID).Find(&articles); dbc.Error != nil {
		return articles, dbc.Error
	}else{
		return articles, nil
	}
}

//Update ...
func (m ArticleModel) Update(userID int64, id int64, form forms.UpdateArticleForm) (err error) {
	//METHOD 1
	//Check the article by ID using this way
	_, err = m.One(userID, id)
	if err != nil {
		return err
	}
	var articleObj Article
	if dbc := db.GetDB().Model(&articleObj).Where(id).Where("user_id",userID).Updates(Article{Title: form.Title, Content: form.Content}); dbc.Error != nil {
		return dbc.Error
	}

	return err
}

//Delete ...
func (m ArticleModel) Delete(userID, id int64) (err error) {

	var articleObj Article
	if dbc := db.GetDB().Delete(&articleObj, id); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}
