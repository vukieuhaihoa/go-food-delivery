package restaurantstorage

import (
	"context"

	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (storage *sqlStorage) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.PrepareForInsert()

	if err := storage.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
