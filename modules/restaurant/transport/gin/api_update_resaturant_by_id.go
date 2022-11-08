package restaurantgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
	"gorm.io/gorm"
)

func UpdateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := restaurantstorage.NewSQLStorage(db)
		business := restaurantbusiness.NewUpdateRestaurantBusiness(storage)

		if err := business.UpdateRestaurantById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}
