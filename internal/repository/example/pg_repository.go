package example

import (
	"context"
	"gorm.io/gorm"
	"orc-system/internal/model"
	"time"
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
	result.ID = 1
	result.Name = "test"
	result.CreatedAt = time.Now()
	result.UpdatedAt = time.Now()

	// xu ly sql
	// err = e.db.Where("id = ?", param.id).First(&result).Error
	return result, err
}

func (e *exampleRepository) GetAllUser(ctx context.Context) ([]model.User, error) {
	var userList []model.User
	err := e.db.Select("*").Table(model.UserTable).Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}
