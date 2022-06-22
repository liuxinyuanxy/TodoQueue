package model

import "time"

type Todo struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   []Subtask
	LastWorkT  time.Time //上一次进行todo的时间
	Ddl        time.Time //ddl
	EstimatedT time.Time //预计耗时
	SpentT     time.Time //实际耗时
	Priority   float64   //优先级
}

type TodoDone struct {
	ID         uint `gorm:"not null;autoIncrement"`
	UID        uint
	Title      string
	Subtasks   []Subtask
	EstimatedT time.Time
	SpentT     time.Time
}

type Subtask struct {
	Title string //子项的标题
	Done  bool   //是否完成
}
