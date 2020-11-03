package models

import "time"

type (
	LogModel struct {
		ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
		UserId    int       `json:"user_id"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func (l *LogModel) Insert() {
	l.CreatedAt = time.Now()
	db.DB.Create(l)
}

func (l *LogModel) FindNot(userId int) ([]LogModel, error) {
	var lm []LogModel
	model := db.DB.Order("id DESC").Where("user_id != ?", userId).Find(&lm)
	return lm, model.Error
}

func (l *LogModel) Find() ([]LogModel, error) {
	var lm []LogModel
	model := db.DB.Order("id DESC").Where(&l).Find(&lm)
	return lm, model.Error
}
