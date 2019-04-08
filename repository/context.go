package repository

import (
	"context"
	"octane/record"
	"strconv"
)

// GetContextsByCategoryID ...
func (l *lottoMySQL) GetContextsByCategoryID(ctx context.Context, categoryID string) ([]record.Context, error) {
	var contexts []record.Context
	err := combineErrors(l.DB.Where("category_id = ?", categoryID).Find(&contexts).GetErrors())
	if err != nil {
		return contexts, err
	}
	return contexts, nil
}

// GetContext ...
func (l *lottoMySQL) GetContext(ctx context.Context, contextId string) (*record.Context, error) {
	var con record.Context
	err := combineErrors(l.DB.Where("id = ?", contextId).First(&con).GetErrors())
	if err == nil {
		return &con, nil
	}
	return nil, err
}

// UpdateContext ...
func (l *lottoMySQL) UpdateContext(ctx context.Context, context *record.Context) error {
	return combineErrors(l.DB.Save(context).GetErrors())
}

// DeleteContext ...
func (l *lottoMySQL) DeleteContext(ctx context.Context, ctxId string) error {
	i, err := strconv.Atoi(ctxId)
	if err != nil {
		return err
	}
	var ct = record.Context{
		ID: int(i),
	}
	return combineErrors(l.DB.Unscoped().Delete(ct).GetErrors())
}

// AddContext ...
func (l *lottoMySQL) AddContext(ctx context.Context, context *record.Context) error {
	return combineErrors(l.DB.Create(context).GetErrors())
}
