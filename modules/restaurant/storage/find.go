package restaurantstorage

import (
	"context"

	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (store *sqlStorage) FindRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := store.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
