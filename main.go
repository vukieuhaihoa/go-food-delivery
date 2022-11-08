package main

import (
	"log"

	"github.com/gin-gonic/gin"
	restaurantgin "github.com/vukieuhaihoa/go-food-delivery/modules/restaurant/transport/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL := "postgresql://root:root@localhost:5432/food_delivery?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

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
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(db))
			restaurants.GET("/:restaurant_id", restaurantgin.FindRestaurantHandler(db))
			restaurants.GET("", restaurantgin.ListRestaurantsHandler(db))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(db))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(db))
		}
	}

	return r.Run()
}

// func createRestaurant(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var data RestaurantCreate
// 		if err := ctx.ShouldBind(&data); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := data.Validate(); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Create(&data).Error; err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data": data.Id,
// 		})

// 	}
// }

// func getRestaurant(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var data Restaurant
// 		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data": data,
// 		})

// 	}
// }

// func getListRestaurants(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var data []Restaurant
// 		type Paging struct {
// 			Page  int   `json:"page" form:"page"`
// 			Limit int   `json:"limit" form:"limit"`
// 			Total int64 `json:"total" form:"-"`
// 		}
// 		var paging Paging
// 		if err := ctx.ShouldBind(&paging); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 		}

// 		if paging.Page <= 0 {
// 			paging.Page = 1
// 		}

// 		if paging.Limit <= 0 {
// 			paging.Limit = 10
// 		}

// 		offset := (paging.Page - 1) * paging.Limit

// 		if err := db.Table(Restaurant{}.TableName()).Count(&paging.Total).Limit(paging.Limit).Offset(offset).Order("id desc").Find(&data).Error; err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data":   data,
// 			"paging": paging,
// 		})

// 	}
// }

// func updateRestaurant(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var data RestaurantUpdate

// 		if err := ctx.ShouldBind(&data); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data": true,
// 		})

// 	}
// }

// func deleteRestaurant(db *gorm.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data": true,
// 		})

// 	}
// }
