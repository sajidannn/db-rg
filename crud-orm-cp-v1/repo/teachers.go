package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	err := t.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	if err := t.db.Select("*").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	if err := t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

func (t TeacherRepo) Delete(id uint) error {
	if err := t.db.Delete(&model.Teacher{}, id).Error; err != nil {
		return err
	}
	return nil
}
