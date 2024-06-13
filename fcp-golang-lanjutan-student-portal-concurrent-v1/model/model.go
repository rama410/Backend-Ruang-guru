package model

type Student struct {
	ID           string
	Name         string
	StudyProgram string
}

type StudentModifier func(*Student) error
