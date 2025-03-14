package api

import (
	"test-plus/internal/interfaces"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogoutHandler_Logout(t *testing.T) {
	type fields struct {
		LogoutService interfaces.ILogoutService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &LogoutHandler{
				LogoutService: tt.fields.LogoutService,
			}
			api.Logout(tt.args.c)
		})
	}
}
