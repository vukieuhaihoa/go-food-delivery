package restaurantstorage

import (
	"context"

	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (storage *sqlStorage) UpdateRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := storage.db
	if err := db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
