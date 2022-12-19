package model

//;;import "gorm.io/gorm"

//User -> model for users table
type User struct {
	ID int     `json:"id"`
	NAME    string `json:"name"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
	CREATEDAT string `json:"createdat"`
	UPDATEDAT string `json:"updatedat"`
}

//TableName --> Table for Users Model
func (User) TableName() string {
	return "users"
}