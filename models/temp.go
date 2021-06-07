package models

import "time"

type File struct {
	Id        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Hash      string    `json:"hash"`
	Size      string    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *File) Insert() {
	t.CreatedAt = time.Now()
	db.DB.Create(t)
}

func (t *File) Find() ([]LogModel, error) {
	var lm []LogModel
	model := db.DB.Order("id DESC").Where(&t).Find(&lm)
	return lm, model.Error
}

func (t *File) Delete() {
	db.DB.Delete(t)
}
