package restaurantstorage

import (
	"context"

	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (storage *sqlStorage) DeleteRestaurant(ctx context.Context, cond map[string]interface{}) error {
	db := storage.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where(cond).
		Delete(nil).
		Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
