package api

import (
	"test-plus/internal/interfaces"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterHandler_Register(t *testing.T) {
	type fields struct {
		RegisterService interfaces.IRegisterService
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
			api := &RegisterHandler{
				RegisterService: tt.fields.RegisterService,
			}
			api.Register(tt.args.c)
		})
	}
}
