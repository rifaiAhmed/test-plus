package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"test-plus/constants"
	"test-plus/helpers"
	"test-plus/internal/mocks"
	models "test-plus/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestTransactionAPI_Create(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockITransactionService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		request    any
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name: "success",
			request: models.Transaction{
				CustomerID:    1,
				NomorKontrak:  "TR2503090001",
				Otr:           7000000,
				JumlahCicilan: 1000000,
				JumlahBunga:   5,
				JumlahBulan:   4,
				NamaAsset:     "Motor",
			},
			mockFn: func() {
				mockService.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(&models.Transaction{
					ID:            1,
					CustomerID:    1,
					NomorKontrak:  "TR2503090001",
					Otr:           7000000,
					JumlahCicilan: 1000000,
					JumlahBunga:   5,
					JumlahBulan:   4,
					NamaAsset:     "Motor",
				}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "error parsing JSON",
			request:    `{"Nik": "123"`,
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name: "error create",
			request: models.Transaction{
				CustomerID:    1,
				NomorKontrak:  "TR2503090001",
				Otr:           7000000,
				JumlahCicilan: 1000000,
				JumlahBunga:   5,
				JumlahBulan:   4,
				NamaAsset:     "Motor",
			},
			mockFn: func() {
				mockService.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   constants.ErrServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &TransactionAPI{
				TransactionService: mockService,
			}
			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("POST", "/transaction/v1/", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = req

			tt.mockFn()

			api.Create(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestTransactionAPI_Find(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockITransactionService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api := &TransactionAPI{
		TransactionService: mockService,
	}
	router.GET("/transaction/v1/:id", api.Find)

	tests := []struct {
		name           string
		mockFn         func()
		paramID        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "success",
			paramID: "1",
			mockFn: func() {
				mockService.EXPECT().
					FindByTranscID(gomock.Any(), 1).
					Return(models.Transaction{
						ID:            1,
						CustomerID:    2,
						NomorKontrak:  "TR2503090001",
						Otr:           7000000,
						JumlahCicilan: 1000000,
						JumlahBunga:   5,
						JumlahBulan:   4,
						NamaAsset:     "Motor",
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"success"`,
		},
		{
			name:           "error - invalid ID format",
			paramID:        "abc",
			mockFn:         func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"message":"data tidak sesuai"`,
		},
		{
			name:    "error - transaction not found",
			paramID: "999",
			mockFn: func() {
				mockService.EXPECT().
					FindByTranscID(gomock.Any(), 999).
					Return(models.Transaction{}, gorm.ErrRecordNotFound)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"Data not found"`,
		},
		{
			name:    "error - database error",
			paramID: "1",
			mockFn: func() {
				mockService.EXPECT().
					FindByTranscID(gomock.Any(), 1).
					Return(models.Transaction{}, errors.New("some database error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"message":"terjadi kesalahan pada server"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			req, _ := http.NewRequest("GET", "/transaction/v1/"+tt.paramID, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)
		})
	}
}
