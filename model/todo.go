package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type Todo struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  `example:"2006-01-02 07:04:05"` //上一次进行todo的时间
	Ddl        string  `example:"2006-01-02 07:04:05"` //ddl
	EstimatedT string  `example:"10"`                  //预计耗时(min)
	SpentT     string  `example:"10"`                  //实际耗时(min)
	Priority   float64 //优先级
}

type TodoDone struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	EstimatedT string
	SpentT     string
}

type Subtask struct {
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done" binding:"required"`
}

type subtaskList []Subtask

type NewTodoReq struct {
	Title      string `binding:"required" example:"test"`
	Subtasks   subtaskList
	Ddl        string  `example:"2006-01-02 07:04:05"`       //ddl
	EstimatedT string  `example:"10"`                        //预计耗时
	Priority   float64 `binding:"oneof=1 2 3 4" example:"4"` //优先级
}

type ChangeTodoInfoReq struct {
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  `example:"2006-01-02 07:04:05"` //上一次进行todo的时间
	Ddl        string  `example:"2006-01-02 07:04:05"` //ddl
	EstimatedT string  `example:"10"`                  //预计耗时(min)
	SpentT     string  `example:"10"`                  //实际耗时(min)
	Priority   float64 //优先级
}

func (p subtaskList) Value() (driver.Value, error) {
	result, err := json.Marshal(p)
	if err != nil && bytes.Compare(result, []byte("[{}]")) == 0 {
		return nil, nil
	}
	return result, err
}

func (p *subtaskList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

func NewTodo(id ...uint) *Todo {
	iD := uint(0)
	if len(id) > 0 {
		iD = id[0]
	}
	return &Todo{
		ID:         iD,
		UID:        0,
		EstimatedT: "",
		SpentT:     "",
	}
}

func CertificateTodo(todoID, uID uint) (flag bool, err error) {
	var uid uint
	queryTodo := db.Model(&Todo{}).Select("uid").Where("id = ?", todoID)
	queryDone := db.Model(&TodoDone{}).Select("uid").Where("id = ?", todoID)
	queryCTX := db.Raw("? UNION ?", queryTodo, queryDone)
	if err = queryCTX.Pluck("uid", &uid).Error; err != nil || uid == 0 {
		return true, err
	}
	return uid == uID, nil
}

func CreateNewTodo(newTodo *Todo) (err error) {
	err = db.Create(newTodo).Error
	return
}

func QueryTodoById(todoID uint) (ret Todo, err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.First(&ret).Error
	return
}

func QueryTodoListByUID(uid uint) (ret []Todo, err error) {
	filter := db.Where("uid = ?", uid)
	err = filter.Find(&ret).Error
	return
}

func UpdateTodoInfo(newTodo *Todo) (err error) {
	err = db.Model(&Todo{ID: newTodo.ID}).Updates(newTodo).Error
	return
}

func DeleteTodoById(todoID uint) (err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.Delete(&Todo{}).Error
	return
}

func QueryDoneById(doneID uint) (ret TodoDone, err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.First(&ret).Error
	return
}

func QueryDoneListByUID(uid uint) (ret []TodoDone, err error) {
	filter := db.Where("uid = ?", uid)
	err = filter.Find(&ret).Error
	return
}

func CreateDone(newDone *TodoDone) (err error) {
	err = db.Create(newDone).Error
	return
}

func DeleteDoneById(doneID uint) (err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.Delete(&TodoDone{}).Error
	return
}