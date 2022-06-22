package model

import "time"

type Todo struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   []Subtask
	LastWorkT  time.Time //上一次进行todo的时间
	Ddl        time.Time //ddl
	EstimatedT int       //预计耗时
	SpentT     int       //实际耗时
	Priority   float64   //优先级
}

type TodoDone struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   []Subtask
	LastWorkT  time.Time //上一次进行todo的时间
	Ddl        time.Time //ddl
	EstimatedT int       //预计耗时
	SpentT     int       //实际耗时
}

type Subtask struct {
	Title string //子项的标题
	Done  bool   //是否完成
}
