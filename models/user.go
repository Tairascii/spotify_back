package models

type User struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Image     string `json:"image"`
	Password  string `json:"password"`
	IsPremium bool   `json:"is_premium"`
}
