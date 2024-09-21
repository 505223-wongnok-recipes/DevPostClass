package item

import (
	"task-api/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}

func (repo Repository) Create(item *model.Item) error {
	return repo.Database.Create(item).Error
}

func (repo Repository) Find(query model.RequestFindItem) ([]model.Item, error) {
	var results []model.Item

	db := repo.Database

	if statuses := query.Statuses; len(statuses) > 0 {
		db = db.Where("status = ?", statuses)
	}

	if err := db.Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

func (repo Repository) FindByID(id uint) (model.Item, error) {
	var result model.Item

	if err := repo.Database.First(&result, id).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (repo Repository) Replace(item model.Item) error {
	return repo.Database.Model(&item).Updates(item).Error
}