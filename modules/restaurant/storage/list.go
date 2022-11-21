package restaurantstorage

import (
	"context"
	"fmt"

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

	if v := paging.FakeCursor; v != "" {
		fmt.Println(v)
		if uid, err := common.DecomposeUIDFromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset(offset)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Count(&paging.Total).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDb(err)
	}

	if len(result) > 0 {
		result[len(result)-1].Mask(true)
		paging.NextCursor = result[len(result)-1].FakeId.String()
	}

	return result, nil
}
