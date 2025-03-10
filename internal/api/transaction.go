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

type TransactionAPI struct {
	TransactionService interfaces.ITransactionService
}

func (api *TransactionAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Transaction
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	data, err := api.TransactionService.CreateTransaction(c, &req)
	if err != nil {
		log.Error("failed to create Transaction: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *TransactionAPI) Find(c *gin.Context) {
	var (
		log = helpers.Logger
		obj models.Transaction
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil || inputID.ID == 0 {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	obj, err = api.TransactionService.FindByTranscID(c, inputID.ID)
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
