package forms


//CreateArticleForm ...
type CreateArticleForm struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

//UpdateArticleForm ...
type UpdateArticleForm struct {
	Title   string `form:"title" json:"title,omitempty"`
	Content string `form:"content" json:"content,omitempty"`
}
