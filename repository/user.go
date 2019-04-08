package repository

import (
	"context"
	"fmt"
	"octane/entities"
	"octane/record"
	"strconv"
)

// GetUsers ...
func (l *lottoMySQL) GetUsers(ctx context.Context, request *entities.GetUsersRequest) ([]record.User, error) {
	var users []record.User
	err := combineErrors(l.DB.Where("id >= ?", fmt.Sprint(request.Id)).Limit(request.Size).Find(&users).GetErrors())
	if err != nil {
		return users, nil
	}
	return users, err
}

// GetUsers ...
func (l *lottoMySQL) GetUser(ctx context.Context, id string) (*record.User, error) {
	var user record.User
	err := combineErrors(l.DB.Where("id = ?", id).First(&user).GetErrors())
	if err == nil {
		return &user, nil
	}
	return nil, err
}

// UpdateUser ...
func (l *lottoMySQL) UpdateUser(ctx context.Context, user *record.User) error {
	return combineErrors(l.DB.Save(user).GetErrors())
}

// DeleteUser ...
func (l *lottoMySQL) DeleteUser(ctx context.Context, id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var user = record.User{
		ID: int64(i),
	}
	return combineErrors(l.DB.Unscoped().Delete(user).GetErrors())
}

// AddUser ...
func (l *lottoMySQL) AddUser(ctx context.Context, user *record.User) error {
	return combineErrors(l.DB.Create(user).GetErrors())
}

func (l *lottoMySQL) UserExists(ctx context.Context, id string) (bool, error) {
	u, err := l.GetUser(ctx, id)
	return err == nil && u != nil, err
}
