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

type CustomerAPI struct {
	CustomerService interfaces.ICustomerService
}

func (api *CustomerAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.CustomerParam
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	data, err := api.CustomerService.CreateCustomer(c, &req)
	if err != nil {
		log.Error("failed to create Customer: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *CustomerAPI) Find(c *gin.Context) {
	var (
		log = helpers.Logger
		obj models.Customer
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil || inputID.ID == 0 {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	obj, err = api.CustomerService.FindByID(c, inputID.ID)
	if err == gorm.ErrRecordNotFound {
		log.Error("not found data")
		helpers.SendResponseHTTP(c, http.StatusOK, constants.ErNotFound, nil)
		return
	}
	if err != nil {
		log.Error("database error : ", err.Error())
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, obj)
}
