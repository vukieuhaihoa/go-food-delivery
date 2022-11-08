package restaurantbusiness

import (
	"context"

	common "github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
)

type ListRestaurantStorage interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	storage ListRestaurantStorage
}

func NewListRestaurantBusiness(storage ListRestaurantStorage) *listRestaurantBusiness {
	return &listRestaurantBusiness{
		storage: storage,
	}
}

func (biz *listRestaurantBusiness) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.storage.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
