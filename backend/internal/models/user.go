package models

type User struct {
	Id         uint   `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
}

type UserSignUp struct {
	FirstName  string `json:"first_name,omitempty" binding:"required,gte=4,lte=50"`
	SecondName string `json:"second_name,omitempty" binding:"required,gte=4,lte=50"`
	Email      string `json:"email,omitempty" binding:"required,gte=4,lte=250"`
	Username   string `json:"username,omitempty" binding:"required,gte=4,lte=50"`
	Password   string `json:"password,omitempty" binding:"required,gte=4,lte=50"`
}

type UserSignIn struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty" binding:"required"`
}

type UserUpdate struct {
	FirstName  string `json:"first_name,omitempty" binding:"required,gte=4,lte=50"`
	SecondName string `json:"second_name,omitempty" binding:"required,gte=4,lte=50"`
	Email      string `json:"email,omitempty" binding:"required,gte=4,lte=250"`
	Username   string `json:"username,omitempty" binding:"required,gte=4,lte=50"`
	Password   string `json:"password,omitempty" binding:"required,gte=4,lte=50"`
}

type UserCtx string

const (
	UserId         UserCtx = "user_id"
	UserFirstName  UserCtx = "user_first_name"
	UserSecondName UserCtx = "user_second_name"
	UserEmail      UserCtx = "user_email"
	UserUsername   UserCtx = "user_username"
	UserPassword   UserCtx = "user_password"
)
