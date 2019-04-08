package controller

import (
	"context"
	"octane/record"
)

func (l *lotto) GetContextsByCategoryID(ctx context.Context, categoryID string) ([]record.Context, error) {
	return l.r.GetContextsByCategoryID(ctx, categoryID)
}

func (l *lotto) GetContext(ctx context.Context, contextId string) (*record.Context, error) {
	return l.r.GetContext(ctx, contextId)
}

func (l *lotto) UpdateContext(ctx context.Context, context *record.Context) error {
	return l.r.UpdateContext(ctx, context)
}

func (l *lotto) DeleteContext(ctx context.Context, ctxId string) error {
	return l.r.DeleteContext(ctx, ctxId)
}

func (l *lotto) AddContext(ctx context.Context, context *record.Context) error {
	return l.r.AddContext(ctx, context)
}
