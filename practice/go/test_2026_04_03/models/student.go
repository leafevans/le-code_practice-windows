package models

type Student struct {
	ID       int
	Number   string
	Password string
	ClassID  int
	Name     string
	Lessons  []*Lesson `gorm:"many2many:lesson_students"`
}

func (Student) TableName() string {
	return "students"
}
