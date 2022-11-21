package main

import (
	"log"

	"github.com/gin-gonic/gin"
	api "github.com/vukieuhaihoa/go-food-delivery/api/v1"
	"github.com/vukieuhaihoa/go-food-delivery/component"
	middleware "github.com/vukieuhaihoa/go-food-delivery/middleware"

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

	// r := gin.New()

	r.Use(middleware.Recover())

	appCtx := component.NewAppContext(db)

	api.MainRoute(r, appCtx)

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
