package controller

import (
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
}

func UpdateReport(c *gin.Context) {
	id := c.Param("id")
	body := UpdateRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var report models.Report
	if result := database.Database.First(&report, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	report.Title = body.Title
	report.Description = body.Description
	report.Approved = body.Approved

	database.Database.Save(&report)

	c.JSON(http.StatusOK, &report)
}
