package daos

import "github.com/jmoiron/sqlx"

type UserDao struct {
	c *sqlx.DB
}

func NewUserDao(c *sqlx.DB) *UserDao {
	return &UserDao{c: c}
}

func (dao *UserDao) SignInUser(login, password string) (string, error) {
	return "", nil
}
