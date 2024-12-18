package handler

import (
	"calc_service/internal/logger"
	"calc_service/internal/model"
	"calc_service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateHandler(c *gin.Context) {
	log := logger.GetLogger()
	var req model.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("Invalid JSON")
		c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid JSON format"})
		return
	}

	result, err := service.CalculateArithmeticExpression(req.Expression)
	if err != nil {
		log.WithError(err).Error("Error evaluating expression")
		c.JSON(http.StatusUnprocessableEntity, model.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.Response{Result: result})
	log.Infof("Successfully evaluated expression: %s = %s", req.Expression, result)
}
