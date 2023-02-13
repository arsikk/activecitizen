package controller

import (
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteReport(c *gin.Context) {
	id := c.Param("id")

	var report models.Report

	if result := database.Database.First(&report, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return

	}

	database.Database.Delete(&report)
	c.JSON(http.StatusOK, "deleted")
}
