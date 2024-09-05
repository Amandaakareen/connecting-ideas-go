package dto

type UserCreate struct {
	Name     string `json:"name" binding:"required,min=20"`
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=8,max=20"`
}
