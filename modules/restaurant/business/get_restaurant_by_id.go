package restaurantbusiness

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

type FindRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBusiness struct {
	storage FindRestaurantStorage
}

func NewFindRestaurantBusiness(storage FindRestaurantStorage) *findRestaurantBusiness {
	return &findRestaurantBusiness{
		storage: storage,
	}
}

func (biz *findRestaurantBusiness) FindRestaurantById(
	ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	data, err := biz.storage.FindRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
