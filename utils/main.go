package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// User repsent user record
type User struct {
	ID   int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(50)"`
}

// Category ...
type Category struct {
	ID   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(100)"`
}

// Context ...
type Context struct {
	ID         int      `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string   `gorm:"type:varchar(255)"`
	UserID     int      `gorm:"column:user_id;index"`
	User       User     `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	CategoryID int      `gorm:"column:category_id;index"`
	Category   Category `gorm:"foreignkey:CategoryID;association_foreignkey:ID"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	db.LogMode(true)
	err := db.Model(User{}).AutoMigrate(User{}).Error
	if err != nil {
		panic(err)
	}
	db.CreateTable(User{})
	db.CreateTable(Category{})
	db.CreateTable(Context{})

	// create user
	user1 := User{Name: "user1"}
	user2 := User{Name: "user2"}
	db.Create(&user1).Create(&user2)

	// create Category
	db.Create(&Category{Name: "cat1"}).Create(&Category{Name: "cat2"})

	// create contexts
	db.Create(&Context{Name: "c1", UserID: 1, CategoryID: 1}).Create(&Context{Name: "c2", UserID: 1, CategoryID: 2}).Create(&Context{Name: "c3", UserID: 2, CategoryID: 1}).Create(&Context{Name: "c4", UserID: 2, CategoryID: 2})

	var res []Context
	db.Joins("JOIN users on users.id = contexts.user_id").
		Joins("JOIN categories on categories.id = contexts.category_id").
		Where("categories.name=?", "cat2").
		Where("users.name=?", "user2").
		Preload("User").
		Preload("Category").
		Find(&res)

	for _, each := range res {
		fmt.Printf("%+v\n", each)
	}

}
