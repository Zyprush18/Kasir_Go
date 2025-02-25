package models

type Messages struct {
	Messages string `json:"messages"`
	Data    []User `json:"data"`
}

type User struct {
	ID int `json:"id"`
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
	EMAIL string `json:"email" gorm:"unique"`
}


type UserRequest struct {
	USERNAME string `json:"username" validate:"required"`
	PASSWORD string `json:"password" validate:"required"`
	EMAIL string `json:"email" validate:"required,email"`
} 