package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

type DeleteRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	DeleteRestaurant(ctx context.Context, cond map[string]interface{}) error
}

type deleteRestaurantBusiness struct {
	storage DeleteRestaurantStorage
}

func NewDeleteRestaurantBusiness(storage DeleteRestaurantStorage) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{
		storage: storage,
	}
}

func (biz *deleteRestaurantBusiness) DeleteRestaurantById(ctx context.Context, id int) error {
	_, err := biz.storage.FindRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.storage.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
