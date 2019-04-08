package controller

import (
	"context"
	"fmt"
	"octane/record"
)

func (l *lotto) GetCategoriesByUserID(ctx context.Context, userId string) ([]record.Category, error) {
	return l.r.GetCategoriesByUserID(ctx, userId)
}

func (l *lotto) GetCategory(ctx context.Context, categoryId string) (*record.Category, error) {
	return l.r.GetCategory(ctx, categoryId)
}

func (l *lotto) UpdateCategory(ctx context.Context, category *record.Category) error {
	exist, err := l.r.UserExists(ctx, fmt.Sprint(category.UserID))
	if !exist || err != nil {
		return fmt.Errorf("User %v does not exist and error = %s", category.UserID, err.Error())
	}
	return l.r.UpdateCategory(ctx, category)
}

func (l *lotto) DeleteCategory(ctx context.Context, categoryId string) error {
	return l.r.DeleteCategory(ctx, categoryId)
}

func (l *lotto) AddCategory(ctx context.Context, category *record.Category) error {
	exist, err := l.r.UserExists(ctx, fmt.Sprint(category.UserID))
	if !exist || err != nil {
		return fmt.Errorf("User %v does not exist and error = %s", category.UserID, err.Error())
	}
	return l.r.AddCategory(ctx, category)
}
