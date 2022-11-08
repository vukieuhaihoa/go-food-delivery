package restaurantgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
	"gorm.io/gorm"
)

func FindRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("restaurant_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := restaurantstorage.NewSQLStorage(db)

		biz := restaurantbusiness.NewFindRestaurantBusiness(storage)

		data, err := biz.FindRestaurantById(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
