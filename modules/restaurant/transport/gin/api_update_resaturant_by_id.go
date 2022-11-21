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

func UpdateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		uid, err := common.DecomposeUIDFromBase58(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDbConnection())
		business := restaurantbusiness.NewUpdateRestaurantBusiness(storage)

		if err := business.UpdateRestaurantById(ctx.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
