package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/john-henry-alfeche/inventory-app-api/controllers"
	"time"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	{
		categories := api.Group("/categories")
		{
			categories.POST("", controllers.CreateCategory)
			categories.GET("", controllers.GetCategories)
			categories.GET("/:category_id", controllers.GetCategoryById)
			categories.PUT("/:category_id", controllers.UpdateCategoryById)
			categories.DELETE("/:category_id", controllers.DeleteCategory)
		}

		products := api.Group("/products")
		{
			products.POST("", controllers.CreateProduct)
			products.GET("", controllers.GetAllProducts)
			products.GET("/:product_id", controllers.GetProductById)
			products.PUT("/:product_id", controllers.UpdateProductById)
			products.DELETE("/:product_id", controllers.DeleteProductById)
		}

		suppliers := api.Group("suppliers")
		{
			suppliers.POST("", controllers.CreateSupplier)
			suppliers.GET("", controllers.GetAllSuppliers)
			suppliers.GET("/:supplier_id", controllers.GetSupplierById)
			suppliers.PUT("/:supplier_id", controllers.UpdateSupplier)
			suppliers.DELETE("/:supplier_id", controllers.DeleteSupplier)
		}
	}

	return r
}
