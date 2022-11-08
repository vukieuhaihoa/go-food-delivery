package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

type UpdateRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBusiness struct {
	storage UpdateRestaurantStorage
}

func NewUpdateRestaurantBusiness(storage UpdateRestaurantStorage) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{
		storage: storage,
	}
}

func (biz *updateRestaurantBusiness) UpdateRestaurantById(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	_, err := biz.storage.FindRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.storage.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
