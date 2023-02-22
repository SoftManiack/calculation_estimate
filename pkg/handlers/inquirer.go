package handler

import (
	dto "calculation-estimate"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInquirer(c *gin.Context) {

	inquirer, err := h.services.GetInquirer()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, inquirer)
}

func (h *Handler) GetEstimate(c *gin.Context) {

	var inquirer dto.Inquirer

	if err := c.BindJSON(&inquirer); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	fmt.Println("handlers")
	fmt.Println(inquirer)

	estimates, err := h.services.СalculationEstimate(inquirer)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, estimates)
}
