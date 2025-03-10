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

func TestCustomerAPI_Create(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockICustomerService(ctrlMock)

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
			request: models.CustomerParam{
				Nik:          "1234567890123456",
				FullName:     "Budi",
				LegalName:    "Budi Santoso",
				TempatLahir:  "Jakarta",
				TanggalLahir: "2002-11-09",
				Gaji:         7000000,
				FotoKtp:      "ktp.jpg",
				FotoSelfi:    "selfie.jpg",
			},
			mockFn: func() {
				mockService.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Return(&models.Customer{
					ID:          1,
					Nik:         "1234567890123456",
					FullName:    "Budi",
					LegalName:   "Budi Santoso",
					TempatLahir: "Jakarta",
					Gaji:        7000000,
					FotoKtp:     "ktp.jpg",
					FotoSelfi:   "selfie.jpg",
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
			name: "error creating customer",
			request: models.CustomerParam{
				Nik:          "1234567890123456",
				FullName:     "Budi",
				LegalName:    "Budi Santoso",
				TempatLahir:  "Jakarta",
				TanggalLahir: "2002-11-09",
				Gaji:         7000000,
				FotoKtp:      "ktp.jpg",
				FotoSelfi:    "selfie.jpg",
			},
			mockFn: func() {
				mockService.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   constants.ErrServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &CustomerAPI{
				CustomerService: mockService,
			}

			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("POST", "/customer/v1/", bytes.NewBuffer(reqBody))
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

func TestCustomerAPI_Find(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockICustomerService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api := &CustomerAPI{
		CustomerService: mockService,
	}
	router.GET("/customer/v1/:id", api.Find)

	tests := []struct {
		name           string
		mockFn         func()
		paramID        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "success - customer found",
			paramID: "1",
			mockFn: func() {
				mockService.EXPECT().
					FindByID(gomock.Any(), 1).
					Return(models.Customer{
						ID:       1,
						FullName: "Budi",
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
			name:    "error - customer not found",
			paramID: "999",
			mockFn: func() {
				mockService.EXPECT().
					FindByID(gomock.Any(), 999).
					Return(models.Customer{}, gorm.ErrRecordNotFound)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"message":"Data not found"`,
		},
		{
			name:    "error - database error",
			paramID: "1",
			mockFn: func() {
				mockService.EXPECT().
					FindByID(gomock.Any(), 1).
					Return(models.Customer{}, errors.New("some database error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `"message":"terjadi kesalahan pada server"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			req, _ := http.NewRequest("GET", "/customer/v1/"+tt.paramID, nil)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)
		})
	}
}
