package mapper

import (
	"errors"
	lpb "octane/grpc/lotto"
	"octane/record"
	"sort"
	"strconv"
)

func CategoriesToGetCategoriesByUserIDResponse(categories []record.Category) (*lpb.Categories, error) {
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].ID < categories[j].ID
	})
	if len(categories) == 0 {
		return nil, errors.New("no categories available")
	}
	var ret lpb.Categories
	for _, u := range categories {
		ret.Categories = append(ret.Categories, CategoryToCategoryResponse(&u))
	}
	return &ret, nil
}

func CategoryToCategoryResponse(cat *record.Category) *lpb.Category {
	return &lpb.Category{
		Id:     int64(cat.ID),
		Name:   cat.Name,
		UserId: int64(cat.UserID),
	}
}

func AddCategoryRequestToCategoryRecord(r *lpb.AddCategoryRequest) *record.Category {
	id, _ := strconv.ParseInt(r.UserId, 10, 32)
	return &record.Category{
		UserID: int(id),
		Name:   r.Name,
	}
}

func UpdateCategoryRequestToCategoryRecord(r *lpb.UpdateCategoryRequest) *record.Category {
	userid, _ := strconv.ParseInt(r.UserId, 10, 32)
	id, _ := strconv.ParseInt(r.Id, 10, 32)
	return &record.Category{
		UserID: int(userid),
		ID:     int(id),
		Name:   r.Name,
	}
}
