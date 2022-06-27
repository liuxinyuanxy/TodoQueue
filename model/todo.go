package model

import (
	"TodoQueue/utils"
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type Todo struct {
	ID         uint `gorm:"not null;autoIncrement;primaryKey"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  `example:"2006-01-02 07:04:05"` //上一次进行todo的时间
	Ddl        string  `example:"2006-01-02 07:04:05"` //ddl
	EstimatedT uint    `example:"10"`                  //预计耗时(min)
	SpentT     uint    `example:"10"`                  //实际耗时(min)
	Priority   float64 //优先级
}

type TodoDone struct {
	ID         uint `gorm:"column:tid;index;not null;primaryKey"`
	UID        uint
	Title      string
	Subtasks   subtaskList
	EstimatedT uint
	SpentT     uint
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
	EstimatedT uint    `example:"10"`                        //预计耗时
	Priority   float64 `binding:"oneof=1 2 3 4" example:"4"` //优先级
}

type GetTodoInfoResp struct {
	ID         uint
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  `example:"2006-01-02 07:04:05"` //上一次进行todo的时间
	Ddl        string  `example:"2006-01-02 07:04:05"` //ddl
	EstimatedT uint    `example:"10"`                  //预计耗时(min)
	SpentT     uint    `example:"10"`                  //实际耗时(min)
	Priority   float64 //优先级
}

type ChangeTodoInfoReq struct {
	Title      string
	Subtasks   subtaskList
	LastWorkT  string  `example:"2006-01-02 07:04:05"` //上一次进行todo的时间
	Ddl        string  `example:"2006-01-02 07:04:05"` //ddl
	EstimatedT uint    `example:"10"`                  //预计耗时(min)
	SpentT     uint    `example:"10"`                  //实际耗时(min)
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
		EstimatedT: 0,
		SpentT:     0,
	}
}

func CertificateTodo(todoID, uID uint) (flag bool, err error) {
	var uid uint
	queryTodo := db.Model(&Todo{}).Select("uid").Where("id = ?", todoID)
	queryDone := db.Model(&TodoDone{}).Select("uid").Where("tid = ?", todoID)
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

func QueryOrderedTodoListByUID(uid uint) (ret []Todo, err error) {
	filter := db.Where("uid = ?", uid)
	filter = filter.Not("priority = 0")
	filter = filter.Order("priority ASC")
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
	filter := db.Where("tid = ?", doneID)
	err = filter.First(&ret).Error
	return
}

func QueryDoneListByUID(uid uint) (ret []TodoDone, err error) {
	filter := db.Where("uid = ?", uid)
	err = filter.Find(&ret).Error
	return
}

func MoveTodo2Done(todoID uint) (err error) {
	var todoRec Todo
	if err = db.First(&todoRec, todoID).Error; err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if er := tx.Delete(&Todo{ID: todoID}).Error; er != nil {
			return er
		}
		if todoRec.Priority == 0 {
			todoRec.SpentT += utils.TimeDuration(todoRec.LastWorkT, utils.CurrentTime())
		}
		newDone := TodoDone{
			ID:         todoRec.ID,
			UID:        todoRec.UID,
			Title:      todoRec.Title,
			Subtasks:   todoRec.Subtasks,
			EstimatedT: todoRec.EstimatedT,
			SpentT:     todoRec.SpentT,
		}
		if er := tx.Create(&newDone).Error; er != nil {
			return er
		}

		return nil
	})
	if err != nil {
		return err
	}
	return
}

func DeleteDoneById(doneID uint) (err error) {
	filter := db.Where("tid = ?", doneID)
	err = filter.Delete(&TodoDone{}).Error
	return
}

func StartTodoById(todoID uint) (err error) {
	updater := map[string]interface{}{"priority": 0, "last_work_t": utils.CurrentTime()}
	result := db.Model(&Todo{ID: todoID}).Updates(updater)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	return
}

func SuspendTodoById(todoID uint) (err error) {
	var inProgress Todo
	if err = db.First(&inProgress, todoID).Error; err != nil {
		return err
	}
	if inProgress.Priority != 0 {
		return gorm.ErrInvalidData
	}

	inProgress.Priority = 1
	inProgress.SpentT += utils.TimeDuration(inProgress.LastWorkT, utils.CurrentTime())

	if err = db.Model(&Todo{ID: todoID}).Updates(&inProgress).Error; err != nil {
		return err
	}
	return
}

func GetTodoInProgressByUid(uid uint) (todoID uint, err error) {
	filter := db.Model(&Todo{}).Where("uid = ?", uid)
	filter = filter.Where("priority = ?", 0)
	if err = filter.Pluck("id", &todoID).Error; err != nil {
		return
	}
	return
}
