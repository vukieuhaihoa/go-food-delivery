package restaurantgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-food-delivery/common"
	"github.com/vukieuhaihoa/go-food-delivery/component"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
)

func CreateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDbConnection())

		business := restaurantbusiness.NewCreateRestaurantBusiness(storage)

		if err := business.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		data.Mask(common.DbTypeRestaurant)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
