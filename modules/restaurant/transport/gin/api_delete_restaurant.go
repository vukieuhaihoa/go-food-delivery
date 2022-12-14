package restaurantgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-food-delivery/common"
	"github.com/vukieuhaihoa/go-food-delivery/component"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
)

func DeleteRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.DecomposeUIDFromBase58(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDbConnection())

		business := restaurantbusiness.NewDeleteRestaurantBusiness(storage)

		if err := business.DeleteRestaurantById(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
