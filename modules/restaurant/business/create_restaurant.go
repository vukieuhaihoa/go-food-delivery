package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

type CreateRestaurantStorage interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	storage CreateRestaurantStorage
}

func NewRestaurantBusiness(storage CreateRestaurantStorage) *createRestaurantBusiness {
	return &createRestaurantBusiness{storage: storage}
}

func (biz *createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.storage.CreateRestaurant(ctx, data); err != nil {
		return err
	}

	return nil
}
