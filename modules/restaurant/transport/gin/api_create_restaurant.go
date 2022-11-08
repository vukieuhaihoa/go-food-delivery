package restaurantgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
	"gorm.io/gorm"
)

func CreateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := restaurantstorage.NewSQLStorage(db)

		business := restaurantbusiness.NewRestaurantBusiness(storage)

		if err := business.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
