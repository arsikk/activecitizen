package controller

import (
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetReports(c *gin.Context) {
	var reports []models.Report

	if result := database.Database.Find(&reports); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &reports)
}
