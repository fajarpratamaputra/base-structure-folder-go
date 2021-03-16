package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// type TodoModel struct {
// 	gorm.Model
// 	Name string `json:"name"`
// }

func (b *User) TableName() string {
	return "user"
}
