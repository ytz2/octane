package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"octane/entities"
	"octane/record"

	"github.com/jinzhu/gorm"
)

const (
	dbstr = "%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local"
)

// LottoMySQL ...
type LottoMySQL interface {
	// Users related
	GetUsers(ctx context.Context, request *entities.GetUsersRequest) ([]record.User, error)
	GetUser(ctx context.Context, id string) (*record.User, error)
	UpdateUser(ctx context.Context, user *record.User) error
	DeleteUser(ctx context.Context, id string) error
	AddUser(ctx context.Context, user *record.User) error
	UserExists(ctx context.Context, id string) (bool, error)

	// Category related
	GetCategoriesByUserID(ctx context.Context, userId string) ([]record.Category, error)
	GetCategory(ctx context.Context, categoryId string) (*record.Category, error)
	UpdateCategory(ctx context.Context, category *record.Category) error
	DeleteCategory(ctx context.Context, categoryId string) error
	AddCategory(ctx context.Context, category *record.Category) error

	// Context related
	GetContextsByCategoryID(ctx context.Context, categoryID string) ([]record.Context, error)
	GetContext(ctx context.Context, contextId string) (*record.Context, error)
	UpdateContext(ctx context.Context, context *record.Context) error
	DeleteContext(ctx context.Context, ctxId string) error
	AddContext(ctx context.Context, context *record.Context) error
}

type lottoMySQL struct {
	DB *gorm.DB
}

// NewLotto ...
func NewLotto(c *record.Config) LottoMySQL {
	if c == nil {
		return nil
	}
	var l lottoMySQL
	con := fmt.Sprintf(dbstr, c.DB.User, c.DB.Password, c.DB.URI, c.DB.Port, c.DB.Name)
	db, err := gorm.Open("mysql", con)
	if err != nil || db == nil {
		log.Fatalf("Failed to open lotto mysql: %s", err)
	}
	l.DB = db
	log.Println("Connected to database lotto")
	return &l
}

func combineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	var errstr = errs[0].Error()
	for i := 1; i < len(errs); i++ {
		errstr += ";" + errs[i].Error()
	}
	return errors.New(errstr)
}
