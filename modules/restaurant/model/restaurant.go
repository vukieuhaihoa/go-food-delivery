package restaurantmodel

import (
	"strings"
	"time"

	"github.com/vukieuhaihoa/go-food-delivery/common"
)

const EntityName = "Restaurant"

var (
	ErrNameCannotBeBlank = common.NewCustomError(nil, "restaurant name can not be blank", "ErrNameCannotBeBlank")
)

type Restaurant struct {
	common.SQLModel
	Owner_id int    `json:"owner_id" gorm:"column:owner_id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Address  string `json:"address" gorm:"column:address;"`
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.SQLModel.Mask(common.DbTypeRestaurant)
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel
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
		return ErrNameCannotBeBlank
	}

	return nil
}
