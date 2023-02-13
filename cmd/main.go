package main

import (
	"ActiveCitizenRESTAPI/controller"
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/middleware"
	"ActiveCitizenRESTAPI/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")
	loadDatabase()
	serverApp()

}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Report{})
}

func serverApp() {
	router := gin.Default()

	auth := router.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())

	auth.POST("/addreport", controller.AddReport)
	auth.PUT("/updatereport", controller.UpdateReport)
	auth.GET("/getreports", controller.GetReports)
	auth.GET("/getreport", controller.GetReport)
	auth.DELETE("/deletereport", controller.DeleteReport)

	//router.POST("/addreport", controller.AddReport)
	//router.PUT("/updatereport", controller.UpdateReport)
	//router.GET("/getreports", controller.GetReports)
	//router.GET("/getreport", controller.GetReport)
	//router.DELETE("/deletereport", controller.DeleteReport)

	publicRoutes := router.Group("/auth")

	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	router.Run(":8000")
	fmt.Println("server runned on :8000")
}
