package restaurantstorage

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (storage *sqlStorage) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := storage.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
