package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	// TODO: replace this

	result := s.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	// TODO: replace this

	result := s.db.Where("token", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Where("username = ?", session.Username).Updates(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	// TODO: replace this
	result := s.db.First(&model.Session{}, "username = ?", name)
	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return result.Error
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	// TODO: replace this
	var session model.Session
	result := s.db.First(&session, "token = ?", token)
	if result.Error != nil {
		return session, result.Error
	}
	return session, nil
}
