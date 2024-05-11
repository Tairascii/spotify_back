package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"spotify_back/models"
)

type UserDao struct {
	c *sqlx.DB
}

func NewUserDao(c *sqlx.DB) *UserDao {
	return &UserDao{c: c}
}

func (dao *UserDao) SignUpUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into users (name, surname, password, image, email, is_premium) values ($1, $2, $3, $4, $5, $6) returning id")
	row := dao.c.QueryRow(query, user.Name, user.Surname, user.Password, "", user.Email, false)

	if err := row.Scan(&id); err != nil {
		log.Printf("something went wrong while creating user %s", err.Error())
		return 0, err
	}
	return id, nil
}

func (dao *UserDao) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id from users where email=$1 and password=$2")
	err := dao.c.Get(&user, query, login, password)
	return user, err
}
