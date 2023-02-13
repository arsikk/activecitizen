package controller

import (
	"ActiveCitizenRESTAPI/database"
	"ActiveCitizenRESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddReportRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
}

func AddReport(c *gin.Context) {
	body := AddReportRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var report models.Report

	report.Title = body.Title
	report.Description = body.Description
	report.Approved = body.Approved

	if result := database.Database.Create(&report); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	}

	c.JSON(http.StatusCreated, &report)
}
