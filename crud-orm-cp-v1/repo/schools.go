package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SchoolRepo struct {
	db *gorm.DB
}

func NewSchoolRepo(db *gorm.DB) SchoolRepo {
	return SchoolRepo{db}
}

func (s SchoolRepo) Init(data []model.School) error {
	err := s.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
