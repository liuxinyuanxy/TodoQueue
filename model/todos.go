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
