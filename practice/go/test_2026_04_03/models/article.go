package models

type Article struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	CateID      int         `json:"cate_id"`
	State       int         `json:"state"`
	ArticleCate ArticleCate `gorm:"foreignKey;references:ID"`
}

func TableName() string {
	return "articles"
}
