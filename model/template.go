package model

type Template struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	EstimatedT uint
	Priority   float64
}

type TemplateRequest struct {
	Title      string      `validate:"required"`
	Subtasks   subtaskList `validate:"required"`
	EstimatedT uint        `validate:"required"`
	Priority   float64     `validate:"required"`
}

type ChangeTemplateReq struct {
	Title      string      `validate:"required"`
	Subtasks   subtaskList `validate:"required"`
	EstimatedT uint        `validate:"required"`
	Priority   float64     `validate:"required"`
}

type GetTemplateInfoResp struct {
	ID         uint        `gorm:"not null;autoIncrement"`
	Title      string      `validate:"required"`
	Subtasks   subtaskList `validate:"required"`
	EstimatedT uint        `validate:"required"`
	Priority   float64     `validate:"required"`
}

func AddTemplate(t *Template) error {
	return db.Create(t).Error
}

func DeleteTemplate(tid uint) error {
	return db.Delete(&Template{ID: tid}).Error
}

func GetAllTemplate(uid uint) (*[]Template, error) {
	var dbTemplates []Template
	result := db.Where(&Template{UID: uid}).Find(&dbTemplates)
	return &dbTemplates, result.Error
}

func UpdateTemplateInfo(newTemplate *Template) (err error) {
	err = db.Model(&Template{ID: newTemplate.ID}).Updates(newTemplate).Error
	return
}

func QueryTemplateById(templateID uint) (ret Template, err error) {
	filter := db.Where("id = ?", templateID)
	err = filter.First(&ret).Error
	return
}
