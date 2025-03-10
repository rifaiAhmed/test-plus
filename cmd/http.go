package cmd

import (
	"log"
	"test-plus/helpers"
	"test-plus/internal/api"
	"test-plus/internal/interfaces"
	"test-plus/internal/repository"
	"test-plus/internal/services"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	// Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// users
	userV1 := r.Group("/user/v1")
	userV1.POST("/register", d.RegisterAPI.Register)
	userV1.POST("/login", d.LoginAPI.Login)

	userV1WithAuth := userV1.Use()
	userV1WithAuth.DELETE("/logout", d.MiddlewareValidateAuth, d.LogoutAPI.Logout)
	userV1WithAuth.PUT("/refresh-token", d.MiddlewareRefreshToken, d.RefreshTokenAPI.RefreshToken)

	// customer
	customerV1 := r.Group("/customer/v1")
	customerV1.POST("/", d.MiddlewareValidateAuth, d.CustomerAPI.Create)
	customerV1.GET("/:id", d.MiddlewareValidateAuth, d.CustomerAPI.Find)

	// limit
	creditLimitV1 := r.Group("/limit/v1")
	creditLimitV1.POST("/", d.MiddlewareValidateAuth, d.CreditLimitAPI.Create)
	creditLimitV1.GET("/:id", d.MiddlewareValidateAuth, d.CreditLimitAPI.Find)

	// transaction
	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/", d.MiddlewareValidateAuth, d.TransactionAPI.Create)
	transactionV1.GET("/:id", d.MiddlewareValidateAuth, d.TransactionAPI.Find)

	err := r.Run(":" + helpers.GetEnv("PORT", "8000"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	RegisterAPI     interfaces.IRegisterHandler
	LoginAPI        interfaces.ILoginHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	CustomerAPI    interfaces.ICustomerAPI
	CreditLimitAPI interfaces.ICreditLimitAPI
	TransactionAPI interfaces.ITransactionAPI
}

func dependencyInject() Dependency {

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		UserRepo: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}
	logoutAPI := &api.LogoutHandler{

		LogoutService: logoutSvc,
	}
	refreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}
	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	// customer
	customerRepo := &repository.CustomerRepo{
		DB: helpers.DB,
	}

	customerSvc := &services.CustomerService{
		CustomerRepo: customerRepo,
	}

	customerAPI := &api.CustomerAPI{
		CustomerService: customerSvc,
	}

	// credit limit
	creditLimitRepo := &repository.CreditLimitRepo{
		DB: helpers.DB,
	}

	creditLimitSvc := &services.CreditLimitService{
		CreditLimitRepo: creditLimitRepo,
	}

	creditLimitAPI := &api.CreditLimitAPI{
		CreditLimitService: creditLimitSvc,
	}

	// transaction
	transactionRepo := &repository.TransactionRepo{
		DB: helpers.DB,
	}

	transactionSvc := &services.TransactionService{
		TransactionRepo: transactionRepo,
	}

	transactionAPI := &api.TransactionAPI{
		TransactionService: transactionSvc,
	}

	return Dependency{
		UserRepository:  userRepo,
		RegisterAPI:     registerAPI,
		LoginAPI:        loginAPI,
		LogoutAPI:       logoutAPI,
		RefreshTokenAPI: refreshTokenAPI,
		// TokenValidationAPI: tokenValidationAPI,
		CustomerAPI:    customerAPI,
		CreditLimitAPI: creditLimitAPI,
		TransactionAPI: transactionAPI,
	}
}
