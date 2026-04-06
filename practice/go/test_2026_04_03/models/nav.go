package models

type Nav struct {
	ID     int
	Title  string
	URL    string
	Status int
	Sort   int
}

func (Nav) TableName() string {
	return "navs"
}
