package example

import (
	"context"
	"gorm.io/gorm"
	"orc-system/internal/model"
)

func NewExampleRepository(db *gorm.DB) IRepository {
	return &exampleRepository{db}
}

type exampleRepository struct {
	db *gorm.DB
}

func (e *exampleRepository) GetExampleByID(param GetExampleByIDInput, ctx context.Context) (model.Example, error) {
	var (
		result model.Example
		err    error
	)
	if err := param.Validate(); err != nil {
		return result, err
	}
	// xu ly sql
	// err = e.db.Where("id = ?", param.id).First(&result).Error
	return result, err
}
