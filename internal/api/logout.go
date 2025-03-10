package api

import (
	"net/http"
	"test-plus/constants"
	"test-plus/helpers"
	"test-plus/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("failed on logout service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}
