package models

type LessonStudent struct {
	LessonID  int
	StudentID int
}

func (LessonStudent) TableName() string {
	return "lesson_students"
}
