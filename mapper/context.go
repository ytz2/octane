package mapper

import (
	"errors"
	lpb "octane/grpc/lotto"
	"octane/record"
	"sort"
)

func GetContextsByCategoryIDResponse(contexts []record.Context) (*lpb.Contexts, error) {
	sort.Slice(contexts, func(i, j int) bool {
		return contexts[i].ID < contexts[j].ID
	})
	if len(contexts) == 0 {
		return nil, errors.New("no context available")
	}
	var ret lpb.Contexts
	for _, u := range contexts {
		ret.Contexts = append(ret.Contexts, ContextToContextResponse(&u))
	}
	return &ret, nil
}

func ContextToContextResponse(context *record.Context) *lpb.Context {
	return &lpb.Context{
		Id:         int64(context.ID),
		Name:       context.Name,
		Note:       context.Note,
		UserID:     int64(context.UserID),
		CategoryID: int64(context.CategoryID),
	}
}

func UpdateContextRequestToContextRecord(u *lpb.UpdateContextRequest) *record.Context {
	return &record.Context{
		ID:         int(u.Id),
		Name:       u.Name,
		Note:       u.Note,
		UserID:     int(u.UserID),
		CategoryID: int(u.CategoryID),
	}
}

func AddContextRequestToContextRecord(u *lpb.AddContextRequest) *record.Context {
	return &record.Context{
		Name:       u.Name,
		Note:       u.Note,
		UserID:     int(u.UserID),
		CategoryID: int(u.CategoryID),
	}
}
