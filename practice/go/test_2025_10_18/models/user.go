package models

type User struct {
	ID       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

func (User) TableName() string {
	return "users"
}
