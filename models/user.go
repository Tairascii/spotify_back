package models

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Image     string `json:"image"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	IsPremium bool   `json:"is_premium"`
}
