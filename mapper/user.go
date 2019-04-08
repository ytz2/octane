package mapper

import (
	"errors"
	"fmt"
	"octane/entities"
	lpb "octane/grpc/lotto"
	"octane/record"
	"sort"
	"strconv"
)

// UserRequestToUserID ...
func UserRequestToUserID(in *lpb.GetUserRequest) string {
	return in.Id
}

// UserRecordToUserResponse ...
func UserRecordToUserResponse(r *record.User) *lpb.User {
	var user lpb.User
	user.Id = r.ID
	user.Name = r.Name
	return &user
}

// UpdateUserRequestToUser ...
func UpdateUserRequestToUser(in *lpb.UpdateUserRequest) *record.User {
	user := new(record.User)
	user.Name = in.Name
	user.ID, _ = strconv.ParseInt(in.Id, 10, 64)
	return user
}

// AddUserRequest ...
func AddUserRequestToUser(in *lpb.AddUserRequest) *record.User {
	user := new(record.User)
	user.Name = in.Name
	return user
}

// DeleteUserRequest ...
func DeleteUserRequestToUserID(in *lpb.DeleteUserRequest) string {
	return in.Id
}

func GetUserRequestToEntities(request *lpb.GetUsersRequest) *entities.GetUsersRequest {
	id, _ := strconv.ParseInt(request.Id, 10, 32)
	return &entities.GetUsersRequest{
		Id:   id,
		Size: int(request.PageSize),
	}
}

func UsersToGetUserResponse(users []record.User) (*lpb.GetUserResponse, error) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	if len(users) == 0 {
		return nil, errors.New("no users available")
	}
	var ret lpb.GetUserResponse
	for _, u := range users {
		ret.Users = append(ret.Users, UserRecordToUserResponse(&u))
	}
	ret.NextId = fmt.Sprint(users[len(users)-1].ID)
	return &ret, nil
}
