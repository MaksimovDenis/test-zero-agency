package handler

import (
	"net/http"
	"strconv"
	zeroAgency "test-zero-agency"
	logger "test-zero-agency/logs"

	"github.com/gin-gonic/gin"
)

// @Summary Get News By ID
// @Security ApiKeyAuth
// @Tags news
// @Description Get News by ID
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} zeroAgency.NewsWithCategories
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router /list/{id} [get]
func (h *Handler) handleGetNewsById(c *gin.Context) {

	logger.Log.Info("Handling Get News By Id")

	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		logger.Log.Info("invailed Id params", err)
		newErrorResponse(c, http.StatusBadRequest, "invailed id params")
	}

	list, err := h.service.NewsWithCategories.GetNewsById(id)
	if err != nil {
		logger.Log.Error("Failed to Get News By ID: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

type GetAllNewsResponse struct {
	Success bool                            `json:"Success"`
	News    []zeroAgency.NewsWithCategories `json:"News"`
}

// @Summary Get All News
// @Security ApiKeyAuth
// @Tags news
// @Description Get List of News
// @Accept json
// @Produce json
// @Success 200 {object} GetAllNewsResponse
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router /list [get]
func (h *Handler) handleGetAllNews(c *gin.Context) {

	logger.Log.Info("Handling Get All News")

	list, err := h.service.NewsWithCategories.GetAllNews()
	if err != nil {
		logger.Log.Error("Failed to Get All News: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := GetAllNewsResponse{
		Success: true,
		News:    list,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update News By ID
// @Security ApiKeyAuth
// @Tags news
// @Description Update News by ID
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} zeroAgency.NewsWithCategories
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router /edit/{id} [put]
func (h *Handler) handleUpdateNewsById(c *gin.Context) {
	logger.Log.Info("Handling Update News By Id")

	var input zeroAgency.NewsWithCategories
	if err := c.BindJSON(&input); err != nil {
		logger.Log.Error("Failed to bind JSON: ", err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid JSON input")
		return
	}

	if input.Id == 0 {
		logger.Log.Error("Invalid news ID")
		newErrorResponse(c, http.StatusBadRequest, "invalid news ID")
		return
	}

	if err := h.service.NewsWithCategories.UpdateNewsById(input); err != nil {
		logger.Log.Error("Failed to Update News By ID: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary Delete News By ID
// @Security ApiKeyAuth
// @Tags news
// @Description Delete News by ID
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200 {object} zeroAgency.NewsWithCategories
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router /edit/{id} [delete]
func (h *Handler) handleDeleteNewsById(c *gin.Context) {

	logger.Log.Info("Handling Delete News By Id")

	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		logger.Log.Info("invailed Id params", err)
		newErrorResponse(c, http.StatusBadRequest, "invailed id params")
	}

	if err := h.service.NewsWithCategories.DeleteNewsById(id); err != nil {
		logger.Log.Error("Failed to Delete News By ID: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
