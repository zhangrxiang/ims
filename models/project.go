package models

import "time"

type ProjectModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PHId      int       `json:"ph_id"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

//多数据查询
func (p *ProjectModel) FindBy() (*[]ProjectModel, error) {
	var projects []ProjectModel
	model := db.DB.Order("id DESC").Where(&p).Find(&projects)
	return &projects, model.Error
}

func (p *ProjectModel) Insert() (*ProjectModel, error) {
	model := db.DB.Create(p)
	return model.Value.(*ProjectModel), model.Error
}

func (p *ProjectModel) Update() (*ProjectModel, error) {
	project := &ProjectModel{}
	model := db.DB.Model(project).Updates(p)
	return model.Value.(*ProjectModel), model.Error
}
