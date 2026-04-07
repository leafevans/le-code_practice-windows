package models

type ArticleCate struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	State    int       `json:"state"`
	Articles []Article `gorm:"foreignKey:CateID"`
}

func (ArticleCate) TableName() string {
	return "article_cates"
}
