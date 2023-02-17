package models

type User struct {
	Id          uint   `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	SecondName  string `json:"second_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}

type UserSignUp struct {
	FirstName   string `json:"first_name,omitempty" binding:"required"`
	SecondName  string `json:"second_name,omitempty" binding:"required"`
	Email       string `json:"email,omitempty" binding:"required"`
	PhoneNumber string `json:"phone_number,omitempty" binding:"required"`
	Username    string `json:"username,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
}

type UserSignIn struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty" binding:"required"`
}

type UserUpdate struct {
	FirstName   string `json:"first_name,omitempty"`
	SecondName  string `json:"second_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}

type UserCtx string

const (
	UserId          UserCtx = "user_id"
	UserFirstName   UserCtx = "user_first_name"
	UserSecondName  UserCtx = "user_second_name"
	UserEmail       UserCtx = "user_email"
	UserPhoneNumber UserCtx = "user_phone_number"
	UserUsername    UserCtx = "user_username"
	UserPassword    UserCtx = "user_password"
)
