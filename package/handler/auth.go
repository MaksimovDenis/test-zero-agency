package handler

import (
	"net/http"
	zeroAgency "test-zero-agency"
	logger "test-zero-agency/logs"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Description  create account
// @Tags auth
// @Accept json
// @Produce json
// @Param input body zeroAgency.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router       /auth/sing-up [post]
func (h *Handler) handleSingUp(c *gin.Context) {

	logger.Log.Info("Handling Sing Up")

	var input zeroAgency.User

	if err := c.BindJSON(&input); err != nil {
		logger.Log.Error("Failed to parse input data: ", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" || input.Password == "" {
		errMsg := "Username and password are required"
		logger.Log.Error(errMsg)
		newErrorResponse(c, http.StatusBadRequest, errMsg)
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		logger.Log.Error("Failed to Create User: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type logInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary LogIn
// @Description  login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body logInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {object} Err
// @Failure 404 {object} Err
// @Failure 500 {object} Err
// @Router       /auth/log-in [post]
func (h *Handler) handleSingIn(c *gin.Context) {

	logger.Log.Info("Handling Log In")

	var input logInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" || input.Password == "" {
		errMsg := "Username and password are required"
		logger.Log.Error(errMsg)
		newErrorResponse(c, http.StatusBadRequest, errMsg)
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		logger.Log.Error("Failed to Generate Token: ", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
