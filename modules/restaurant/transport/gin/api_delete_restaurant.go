package restaurantgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
	"gorm.io/gorm"
)

func DeleteRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := restaurantstorage.NewSQLStorage(db)

		business := restaurantbusiness.NewDeleteRestaurantBusiness(storage)

		if err := business.DeleteRestaurantById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}
