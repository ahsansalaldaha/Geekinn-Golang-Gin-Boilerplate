package models

import (
	"errors"

	"github.com/Geekinn/go-micro/db"
	"github.com/Geekinn/go-micro/app/forms"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User ...
type User struct {
	gorm.Model
	ID        int64    `gorm:"primaryKey"`
	Email     string
	Password  string
	Name      string
	UpdatedAt int64    `gorm:"autoUpdateTime"`
	CreatedAt int64	   `gorm:"autoCreateTime"`
}

//UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

func (m UserModel) Migrate(){
	db.GetDB().AutoMigrate(&User{})
}

//Login ...
func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {

	if dbc := db.GetDB().Where("email = ?", form.Email).First(&user); dbc.Error != nil {
		return user, token, dbc.Error
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

//Register ...
func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	
	var exists bool
	err = db.GetDB().Model(user).
			Select("count(*) > 0").
			Where("email = ?", form.Email).
			Find(&exists).
			Error
	
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	if exists {
		return user, errors.New("Email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	userObj := User{Email: form.Email, Password: string(hashedPassword), Name: form.Name}
	if dbc := db.GetDB().Create(&userObj); dbc.Error != nil {
		return user, dbc.Error
	}else{
		return userObj, dbc.Error
	}
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	
	if dbc := db.GetDB().Where(userID).First(&user); dbc.Error != nil {
		return user, dbc.Error
	}
	return user, err
}

//One by email...
func (m UserModel) OneByEmail(email string) (user User, err error) {
	
	if dbc := db.GetDB().Where("email = ?", email).First(&user); dbc.Error != nil {
		return user, dbc.Error
	}
	return user, err
}
