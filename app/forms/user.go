package forms

//LoginForm ...
type LoginForm struct {
	Email    string `form:"email" json:"email" `
	Password string `form:"password" json:"password" `
}

//RegisterForm ...
type RegisterForm struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password" `
}