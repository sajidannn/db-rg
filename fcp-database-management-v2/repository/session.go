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
	err := s.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	err := s.db.Where("token = ?", token).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil
}

// emang gini?
func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	err := s.db.Where("username = ?", session.Username).Updates(&session).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	err := s.db.Where("username = ?", name).First(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var tokenSession model.Session
	err := s.db.Where("token = ?", token).First(&tokenSession).Error
	if err != nil {
		return model.Session{}, err
	}
	return tokenSession, nil
}
