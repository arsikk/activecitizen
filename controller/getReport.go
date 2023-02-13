package controller

import (
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetReport(c *gin.Context) {
	id := c.Param("id")

	var report models.Report
	if result := database.Database.First(&report, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &report)
}
