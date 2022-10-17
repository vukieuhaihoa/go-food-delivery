package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
			restaurants.POST("", createRestaurant(db))
			restaurants.GET("/:restaurant_id", getRestaurant(db))
			restaurants.GET("", getListRestaurants(db))
			restaurants.PUT("/:restaurant_id", updateRestaurant(db))
			restaurants.DELETE("/:restaurant_id", deleteRestaurant(db))
		}
	}

	return r.Run()
}

func createRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data RestaurantCreate
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := data.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
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

func getRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data Restaurant
		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func getListRestaurants(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data []Restaurant
		type Paging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}
		var paging Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		if err := db.Table(Restaurant{}.TableName()).Count(&paging.Total).Limit(paging.Limit).Offset(offset).Order("id desc").Find(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
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

func updateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data RestaurantUpdate

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

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}

func deleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("restaurant_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name       *string    `json:"name" gorm:"column:name;"`
	Address    *string    `json:"address" gorm:"column:address;"`
	Updated_at *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}
