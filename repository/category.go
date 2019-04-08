package repository

import (
	"context"
	"octane/record"
	"strconv"
)

// GetCategoriesByUserID ...
func (l *lottoMySQL) GetCategoriesByUserID(ctx context.Context, userId string) ([]record.Category, error) {
	var categories []record.Category
	err := combineErrors(l.DB.Where("user_id = ?", userId).Find(&categories).GetErrors())
	if err != nil {
		return categories, err
	}
	return categories, nil
}

// GetCategory ...
func (l *lottoMySQL) GetCategory(ctx context.Context, categoryId string) (*record.Category, error) {
	var category record.Category
	err := combineErrors(l.DB.Where("id = ?", categoryId).First(&category).GetErrors())
	if err == nil {
		return &category, nil
	}
	return nil, err
}

// UpdateCategory ...
func (l *lottoMySQL) UpdateCategory(ctx context.Context, category *record.Category) error {
	return combineErrors(l.DB.Save(category).GetErrors())
}

// DeleteCategory ...
func (l *lottoMySQL) DeleteCategory(ctx context.Context, categoryId string) error {
	i, err := strconv.Atoi(categoryId)
	if err != nil {
		return err
	}
	var category = record.Category{
		ID: int(i),
	}
	return combineErrors(l.DB.Unscoped().Delete(category).GetErrors())
}

// AddCategory ...
func (l *lottoMySQL) AddCategory(ctx context.Context, category *record.Category) error {
	return combineErrors(l.DB.Create(category).GetErrors())
}
