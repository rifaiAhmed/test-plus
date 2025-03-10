package api

import (
	"test-plus/internal/interfaces"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginHandler_Login(t *testing.T) {
	type fields struct {
		LoginService interfaces.ILoginService
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
			api := &LoginHandler{
				LoginService: tt.fields.LoginService,
			}
			api.Login(tt.args.c)
		})
	}
}
