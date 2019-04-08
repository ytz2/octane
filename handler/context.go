package handler

import (
	"context"
	lpb "octane/grpc/lotto"
	"octane/mapper"
)

func (l *lotto) GetContextsByCategoryID(ctx context.Context, in *lpb.GetContextsByCategoryIDRequest) (*lpb.Contexts, error) {
	if !l.r.Allow("GetContextsByCategoryID") {
		return nil, limitError
	}
	cs, err := l.c.GetContextsByCategoryID(ctx, in.CategoryID)
	if err != nil {
		return nil, err
	}
	return mapper.GetContextsByCategoryIDResponse(cs)
}

func (l *lotto) GetContext(ctx context.Context, in *lpb.GetContextRequest) (*lpb.Context, error) {
	if !l.r.Allow("GetContext") {
		return nil, limitError
	}
	c, err := l.c.GetContext(ctx, in.ContextID)
	if err != nil {
		return nil, err
	}
	return mapper.ContextToContextResponse(c), nil
}

func (l *lotto) UpdateContext(ctx context.Context, in *lpb.UpdateContextRequest) (*lpb.Void, error) {
	if !l.r.Allow("UpdateContext") {
		return nil, limitError
	}
	err := l.c.UpdateContext(ctx, mapper.UpdateContextRequestToContextRecord(in))
	if err != nil {
		return nil, err
	}
	return &lpb.Void{}, nil
}

func (l *lotto) DeleteContext(ctx context.Context, in *lpb.DeleteContextRequest) (*lpb.Void, error) {
	if !l.r.Allow("DeleteContext") {
		return nil, limitError
	}
	return &lpb.Void{}, l.c.DeleteContext(ctx, in.ContextID)
}

func (l *lotto) AddContext(ctx context.Context, in *lpb.AddContextRequest) (*lpb.Void, error) {
	if !l.r.Allow("AddContext") {
		return nil, limitError
	}
	err := l.c.AddContext(ctx, mapper.AddContextRequestToContextRecord(in))
	if err != nil {
		return nil, err
	}
	return &lpb.Void{}, nil
}
