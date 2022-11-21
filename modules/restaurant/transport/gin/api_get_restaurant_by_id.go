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

func FindRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid, err := common.DecomposeUIDFromBase58(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStorage(appCtx.GetMainDbConnection())

		biz := restaurantbusiness.NewFindRestaurantBusiness(storage)

		data, err := biz.FindRestaurantById(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrCannotGetEntity(restaurantmodel.EntityName, err))
			return
		}

		data.Mask(true)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
