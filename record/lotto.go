package record

// User repsent user record
type User struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"type:varchar(50)"`
}

// TableName ...
func (u User) TableName() string {
	return "user"
}

// Category ...
type Category struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(100)"`
	UserID int    `gorm:"column:user_id;index"`
}

// TableName ...
func (u Category) TableName() string {
	return "category"
}

// Context ...
type Context struct {
	ID         int    `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(255)"`
	Note       string
	UserID     int `gorm:"column:user_id;index"`
	CategoryID int `gorm:"column:category_id;index"`
}

// TableName ...
func (u Context) TableName() string {
	return "context"
}
