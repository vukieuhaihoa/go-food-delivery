package restaurantgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-food-delivery/common"
	"github.com/vukieuhaihoa/go-food-delivery/component"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
)

func ListRestaurantsHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data []restaurantmodel.Restaurant

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			// ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		fmt.Printf("%#v\n", paging)

		_ = paging.Validate()

		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDbConnection())

		business := restaurantbusiness.NewListRestaurantBusiness(storage)

		data, err := business.ListRestaurant(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		for i := range data {
			data[i].Mask(true)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
