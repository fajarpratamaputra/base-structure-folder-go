package Models

import (
	"mvc-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllBooks(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
