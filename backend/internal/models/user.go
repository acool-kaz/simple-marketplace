package models

type User struct {
	Id          uint   `json:"id,omitempty"`
	Role        string `json:"role,omitempty"`
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

var UserSortParams = []string{"id", "role", "first_name", "second_name", "phone_number", "email", "username", "password"}

type UserCtx string

const (
	UserId          UserCtx = "user_id"
	UserRole        UserCtx = "user_role"
	UserFirstName   UserCtx = "user_first_name"
	UserSecondName  UserCtx = "user_second_name"
	UserEmail       UserCtx = "user_email"
	UserPhoneNumber UserCtx = "user_phone_number"
	UserUsername    UserCtx = "user_username"
	UserPassword    UserCtx = "user_password"

	UserSortBy   UserCtx = "user_sort_by"
	UserFilterBy UserCtx = "user_filter_by"
)

const (
	AdminRoleInfo string = "admin"
	UserRoleInfo  string = "user"
)
