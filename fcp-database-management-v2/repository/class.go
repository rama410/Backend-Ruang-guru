package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	// TODO: replace this
	var class []model.Class
	rows, err := s.db.Table("classes").Select("*").Rows()

	for rows.Next() {
		s.db.ScanRows(rows, &class)
	}
	// err := s.db.Find(&class).Error
	return class, err
}
