package handler

import (
	"context"
	lpb "octane/grpc/lotto"
	"octane/mapper"
)

func (l *lotto) GetUser(ctx context.Context, in *lpb.GetUserRequest) (*lpb.User, error) {
	if !l.r.Allow("GetUser") {
		return nil, limitError
	}
	r, err := l.c.GetUser(ctx, mapper.UserRequestToUserID(in))
	if err != nil {
		return nil, err
	}
	ret := mapper.UserRecordToUserResponse(r)
	return ret, nil
}

func (l *lotto) UpdateUser(ctx context.Context, in *lpb.UpdateUserRequest) (*lpb.Void, error) {
	if !l.r.Allow("UpdateUser") {
		return nil, limitError
	}
	ret := new(lpb.Void)
	u := mapper.UpdateUserRequestToUser(in)
	err := l.c.UpdateUser(ctx, u)
	return ret, err
}
func (l *lotto) AddUser(ctx context.Context, in *lpb.AddUserRequest) (*lpb.Void, error) {
	if !l.r.Allow("AddUser") {
		return nil, limitError
	}
	ret := new(lpb.Void)
	u := mapper.AddUserRequestToUser(in)
	err := l.c.AddUser(ctx, u)
	return ret, err
}

func (l *lotto) DeleteUser(ctx context.Context, in *lpb.DeleteUserRequest) (*lpb.Void, error) {
	if !l.r.Allow("DeleteUser") {
		return nil, limitError
	}
	ret := new(lpb.Void)
	id := mapper.DeleteUserRequestToUserID(in)
	err := l.c.DeleteUser(ctx, id)
	return ret, err
}

func (l *lotto) GetUsers(ctx context.Context, in *lpb.GetUsersRequest) (*lpb.GetUserResponse, error) {
	if !l.r.Allow("GetUsers") {
		return nil, limitError
	}
	r := mapper.GetUserRequestToEntities(in)
	ret, err := l.c.GetUsers(ctx, r)
	if err != nil {
		return nil, err
	}
	return mapper.UsersToGetUserResponse(ret)
}
