package models

type ResourceHistoryModel struct {
	ID         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ResourceID int    `json:"resource_id" gorm:"not null"`
	File       string `json:"file"`
	Path       string `json:"path"`
	Hash       string `json:"hash"`
	Version    string `json:"version" gorm:"not null"`
}

func (rh *ResourceHistoryModel) Insert() (*ResourceHistoryModel, error) {
	model := db.DB.Create(rh)
	return model.Value.(*ResourceHistoryModel), model.Error
}

func (rh *ResourceHistoryModel) FindByResourceId() ([]ResourceHistoryModel, error) {
	var resources []ResourceHistoryModel
	model := db.DB.Order("id DESC").Find(&resources, "resource_id = ?", rh.ResourceID)
	return resources, model.Error
}
