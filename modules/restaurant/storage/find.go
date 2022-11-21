package restaurantstorage

import (
	"context"

	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	"gorm.io/gorm"
)

func (store *sqlStorage) FindRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := store.db.Where(cond).First(&data).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
