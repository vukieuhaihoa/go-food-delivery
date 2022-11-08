package restaurantstorage

import (
	"context"

	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

func (store *sqlStorage) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	offset := (paging.Page - 1) * paging.Limit

	db := store.db

	var result []restaurantmodel.Restaurant

	if v := filter.Owner_id; v > 0 {
		db = db.Where("owner_id = ?", v)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Count(&paging.Total).
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}

	return result, nil
}
