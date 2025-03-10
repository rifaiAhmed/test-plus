package api

import (
	"bytes"
	"encoding/json"
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
)

func TestCreditLimitAPI_Create(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockICreditLimitService(ctrlMock)

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
			request: models.CreditLimit{
				CustomerID: 2,
				Tenor1:     100000,
				Tenor2:     200000,
				Tenor3:     300000,
				Tenor4:     400000,
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
			mockFn: func() {
				mockService.EXPECT().CreateCreditLimit(gomock.Any(), gomock.Any()).Return(&models.CreditLimit{
					ID:         1,
					CustomerID: 2,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
				}, nil)
			},
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
			request: models.CreditLimit{
				CustomerID: 2,
				Tenor1:     100000,
				Tenor2:     200000,
				Tenor3:     300000,
				Tenor4:     400000,
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   constants.ErrServerError,
			mockFn: func() {
				mockService.EXPECT().CreateCreditLimit(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &CreditLimitAPI{
				CreditLimitService: mockService,
			}
			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("POST", "/limit/v1/", bytes.NewBuffer(reqBody))
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
