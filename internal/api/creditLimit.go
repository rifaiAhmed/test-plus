package api

import (
	"net/http"
	"test-plus/constants"
	"test-plus/helpers"
	"test-plus/internal/interfaces"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreditLimitAPI struct {
	CreditLimitService interfaces.ICreditLimitService
}

func (api *CreditLimitAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.CreditLimit
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	data, err := api.CreditLimitService.CreateCreditLimit(c, &req)
	if err != nil {
		log.Error("failed to create CreditLimit: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *CreditLimitAPI) Find(c *gin.Context) {
	var (
		log = helpers.Logger
		obj models.CreditLimit
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil || inputID.ID == 0 {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	obj, err = api.CreditLimitService.FindLimitByID(c, inputID.ID)
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error("failed to upadte CreditLimit: ", err)
		helpers.SendResponseHTTP(c, http.StatusOK, "Data not found", nil)
		return
	}
	if err != nil {
		log.Error("failed to upadte CreditLimit: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, obj)
}
