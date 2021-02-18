package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type UserLogin struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
