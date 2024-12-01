package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)

	FetchByID(id int) (*model.Session, error)
}

type sessionsRepoImpl struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (u *sessionsRepoImpl) AddSessions(session model.Session) error {
	_, err := u.db.Exec("INSERT INTO sessions (token, username, expiry) VALUES ($1, $2, $3)", session.Token, session.Username, session.Expiry)
	if err != nil {
		return err
	}

	return nil
}

func (u *sessionsRepoImpl) DeleteSession(token string) error {
	_, err := u.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		return err
	}

	return nil
}

func (u *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	_, err := u.db.Exec("UPDATE sessions SET token = $1, username = $2, expiry = $3 WHERE username = $2", session.Token, session.Username, session.Expiry)
	if err != nil {
		return err
	}
	
	return nil
}

func (u *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	err := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE username = $1", name).Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return err
	}

	return nil	
}

func (u *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE token = $1", token).Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return model.Session{}, err
	}
	
	return session, nil
}

func (u *sessionsRepoImpl) FetchByID(id int) (*model.Session, error) {
	row := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE id = $1", id)

	var session model.Session
	err := row.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
