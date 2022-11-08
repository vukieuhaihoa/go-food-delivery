package restaurantmodel

import (
	"errors"
	"strings"
	"time"

	"github.com/vukieuhaihoa/go-food-delivery/common"
)

type Restaurant struct {
	common.SQLModel
	Owner_id int    `json:"owner_id" gorm:"column:owner_id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Address  string `json:"address" gorm:"column:address;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name       *string    `json:"name" gorm:"column:name;"`
	Address    *string    `json:"address" gorm:"column:address;"`
	Updated_at *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}
