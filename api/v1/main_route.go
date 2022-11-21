package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vukieuhaihoa/go-food-delivery/component"
	restaurantgin "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/transport/gin"
)

func MainRoute(r *gin.Engine, appCtx component.AppContext) {
	v1 := r.Group("/v1")
	{
		healthCheck := v1.Group("/healthcheck")
		{
			healthCheck.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "health check success.",
				})
			})
		}

		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.ListRestaurantsHandler(appCtx))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}
	}
}
