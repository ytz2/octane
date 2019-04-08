package controller

import (
	"context"
	"octane/entities"
	"octane/record"
)

func (l *lotto) GetUser(ctx context.Context, id string) (*record.User, error) {
	return l.r.GetUser(ctx, id)
}

func (l *lotto) GetUsers(ctx context.Context, request *entities.GetUsersRequest) ([]record.User, error) {
	return l.r.GetUsers(ctx, request)
}
func (l *lotto) UpdateUser(ctx context.Context, user *record.User) error {
	return l.r.UpdateUser(ctx, user)
}

func (l *lotto) DeleteUser(ctx context.Context, id string) error {
	return l.r.DeleteUser(ctx, id)
}

func (l *lotto) AddUser(ctx context.Context, user *record.User) error {
	return l.r.AddUser(ctx, user)
}
