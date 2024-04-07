package cars

import (
	"gorm.io/gorm"
	"time"
)

type Car struct {
	ID        int            `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	RegNum    string         `gorm:"column:reg_num"  json:"regNum"`
	Mark      string         `gorm:"column:mark"  json:"mark"`
	Model     string         `gorm:"column:model"  json:"model"`
	Year      int            `gorm:"column:year"  json:"year"`
	OwnerID   int            `gorm:"column:owner_id"  json:"owner_id"`
	CreatedAt *time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Owner     People         `gorm:"foreignkey:OwnerID" json:"-"`
}

func (*Car) TableName() string {
	return "cars"
}

type People struct {
	ID      int    `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Name    string `gorm:"column:name"  json:"name"`
	Surname string `gorm:"column:surname"  json:"surname"`
}

func (*People) TableName() string {
	return "people"
}

type AddCarsRequest struct {
	RegNums []string `json:"regNums"`
}

type ApiResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

type GetCarsFilter struct {
	RegNum    *string `form:"regNum"`
	Mark      *string `form:"mark"`
	Model     *string `form:"model"`
	Year      *int    `form:"year"`
	OwnerID   *int    `form:"ownerID"`
	Page      *int    `form:"page"`
	PageLimit *int    `form:"page_limit"`
}

type Response struct {
	ID        int    `json:"id"`
	RegNum    string `json:"reg_num"`
	Mark      string `json:"mark"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	OwnerID   int    `json:"owner_id"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"created_at"`
}

type GetCarsResponse struct {
	TotalQuantity int        `json:"total_quantity"`
	Page          int        `json:"page"`
	Pages         int        `json:"pages"`
	Cars          []Response `json:"cars"`
}
