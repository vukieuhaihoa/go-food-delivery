package restaurantgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-food-delivery/common"
	restaurantbusiness "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/business"
	restaurantmodel "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/model"
	restaurantstorage "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/storage"
	"gorm.io/gorm"
)

func ListRestaurantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data []restaurantmodel.Restaurant

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		_ = paging.Validate()

		var filter restaurantmodel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		storage := restaurantstorage.NewSQLStorage(db)

		business := restaurantbusiness.NewListRestaurantBusiness(storage)

		data, err := business.ListRestaurant(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":   data,
			"paging": paging,
		})

	}
}
