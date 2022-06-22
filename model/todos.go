package model

type Todo struct {
	ID    uint `gorm:"not null;autoIncrement"`
	UID   uint
	Title string
	//Subtasks   []Subtask
	LastWorkT  string  //上一次进行todo的时间
	Ddl        string  //ddl
	EstimatedT string  //预计耗时
	SpentT     string  //实际耗时
	Priority   float64 //优先级
}

type TodoDone struct {
	ID    uint `gorm:"not null;autoIncrement"`
	UID   uint
	Title string
	//Subtasks   []Subtask
	EstimatedT string
	SpentT     string
}

type Subtask struct {
	Title string //子项的标题
	Done  bool   //是否完成
}
