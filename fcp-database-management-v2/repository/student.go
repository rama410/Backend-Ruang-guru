package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	// TODO: replace this
	var student []model.Student
	result := s.db.Find(&student)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	// TODO: replace this
	result := s.db.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	// TODO: replace this
	result := s.db.Where("id = ?", id).Updates(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	// TODO: replace this

	result := s.db.Delete(&model.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	// TODO: replace this
	var student model.Student
	result := s.db.First(&student, "id = ?", id)
	if result.Error != nil {
		return &student, result.Error
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	// TODO: replace this

	var studentClass []model.StudentClass

	err := s.db.Table("students").Select("students.name as name, students.address as address, classes.name as class_name, classes.professor as professor, classes.room_number as room_number").Joins("JOIN classes ON students.class_id = classes.id").Scan(&studentClass).Error
	if studentClass == nil {
		return &[]model.StudentClass{}, err
	}
	return &studentClass, err
}
