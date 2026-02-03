package routes

import (
	"hrapi/internal/handlers"
	"hrapi/internal/repositories"
	"hrapi/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	regionRepo := repositories.NewRegionRepository(db)
	regionService := services.NewRegionService(regionRepo)
	regionHandler := handlers.NewRegionHandler(regionService)

	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	basePath := viper.GetString("SERVER.BASE_PATH")
	api := router.Group(basePath)
	{
		// region routes endpoints
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.GET("/countries", regionHandler.GetAllRegionsWithCountry)
			regions.GET("/:id/countries", regionHandler.GetRegionByIDWithCountry)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}

		employees := api.Group("/employee")
		{
			employees.GET("", employeeHandler.GetAllEmployees)
			employees.GET("/search", employeeHandler.SearchEmployees)
			employees.POST("", employeeHandler.CreateEmployee)
			employees.PUT("/:id", employeeHandler.UpdateEmployee)
			employees.DELETE("/:id", employeeHandler.DeleteEmployee)
		}
	}
}
