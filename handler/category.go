package handler

import (
	"context"
	lpb "octane/grpc/lotto"
	"octane/mapper"
)

func (l *lotto) GetCategoriesByUserID(ctx context.Context, in *lpb.GetCategoriesByUserIDRequest) (*lpb.Categories, error) {
	if !l.r.Allow("GetCategoriesByUserID") {
		return nil, limitError
	}

	cs, err := l.c.GetCategoriesByUserID(ctx, in.UserID)
	if err != nil {
		return nil, err
	}

	return mapper.CategoriesToGetCategoriesByUserIDResponse(cs)
}

func (l *lotto) GetCategory(ctx context.Context, in *lpb.GetCategoryRequest) (*lpb.Category, error) {
	if !l.r.Allow("GetCategory") {
		return nil, limitError
	}
	c, err := l.c.GetCategory(ctx, in.CategoryID)
	if err != nil {
		return nil, err
	}
	return mapper.CategoryToCategoryResponse(c), nil
}

func (l *lotto) AddCategory(ctx context.Context, in *lpb.AddCategoryRequest) (*lpb.Void, error) {
	if !l.r.Allow("AddCategory") {
		return nil, limitError
	}
	return &lpb.Void{}, l.c.AddCategory(ctx, mapper.AddCategoryRequestToCategoryRecord(in))
}

func (l *lotto) UpdateCategory(ctx context.Context, in *lpb.UpdateCategoryRequest) (*lpb.Void, error) {
	if !l.r.Allow("UpdateCategory") {
		return nil, limitError
	}
	return &lpb.Void{}, l.c.UpdateCategory(ctx, mapper.UpdateCategoryRequestToCategoryRecord(in))
}

func (l *lotto) DeleteCategory(ctx context.Context, in *lpb.DeleteCategoryRequest) (*lpb.Void, error) {
	if !l.r.Allow("DeleteCategory") {
		return nil, limitError
	}
	return &lpb.Void{}, l.c.DeleteCategory(ctx, in.Id)
}
