package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Todo struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  //上一次进行todo的时间
	Ddl        string  //ddl
	EstimatedT string  //预计耗时
	SpentT     string  //实际耗时
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
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type subtaskList []Subtask

func (p subtaskList) Value() (driver.Value, error) {
	return json.Marshal(p)
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
		EstimatedT: 0,
		SpentT:     0,
	}
}

func CreateNewTodo(newTodo *Todo) (err error) {
	err = db.Create(newTodo).Error
	return
}

func QueryTodoById(todoID uint) (ret *Todo, err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.Find(ret).Error
	return
}

func UpdateTodoInfo(newTodo *Todo) (err error) {
	err = db.Model(&Todo{}).Updates(newTodo).Error
	return
}

func DeleteTodoById(todoID uint) (err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.Delete(&Todo{}).Error
	return
}

func QueryDoneById(doneID uint) (ret *TodoDone, err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.Find(ret).Error
	return
}

func DeleteDoneById(doneID uint) (err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.Delete(&TodoDone{}).Error
	return
}


func NewTodo(id ...uint) *Todo {
	var iD uint
	iD = 0
	if len(id) > 0 {
		iD = id[0]
	}
	return &Todo{
		ID:         iD,
		UID:        0,
		EstimatedT: 0,
		SpentT:     0,
	}
}

func CreateNewTodo(newTodo *Todo) (err error) {
	err = db.Create(newTodo).Error
	return
}

func QueryTodoById(todoID uint) (ret *Todo, err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.Find(ret).Error
	return
}

func UpdateTodoInfo(newTodo *Todo) (err error) {
	err = db.Model(&Todo{}).Updates(newTodo).Error
	return
}

func DeleteTodoById(todoID uint) (err error) {
	filter := db.Where("id = ?", todoID)
	err = filter.Delete(&Todo{}).Error
	return
}

func QueryDoneById(doneID uint) (ret *TodoDone, err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.Find(ret).Error
	return
}

func DeleteDoneById(doneID uint) (err error) {
	filter := db.Where("id = ?", doneID)
	err = filter.Delete(&TodoDone{}).Error
	return
}
