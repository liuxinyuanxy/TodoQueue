package model

import "time"

type Template struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   []Subtask
	EstimatedT time.Time
	Priority   float64
}

type TemplateRequest struct {
	Title      string    `validate:"required"`
	Subtasks   []Subtask `validate:"required"`
	EstimatedT time.Time `validate:"required"`
	Priority   float64   `validate:"required"`
}

func AddTemplate(t *Template) error {
	return db.Create(t).Error
}

func DeleteTemplate(tid uint) error {
	return db.Model(&Template{}).Delete(tid).Error
}

func GetAllTemplate(uid uint) (*[]Template, error) {
	var dbTemplates []Template
	result := db.Where(&Template{UID: uid}).Find(&dbTemplates)
	return &dbTemplates, result.Error
}
