package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetupRouter(handler *api.HTTPHandler, userService ports.UserService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{

		r.GET("ping", handler.PingHandler)
		r.POST("/user/beneficiarysignup", handler.FoodBeneficiarySignUp)
		r.PATCH("/user/beneficiaryverifyemail/:token", handler.BeneficiaryVerifyEmail)
		r.POST("/user/kitchenstaffsignup", handler.KitchenStaffSignUp)
		r.PATCH("/user/kitchenstaffverifyemail/:token", handler.KitchenStaffVerifyEmail)
		r.POST("/user/adminsignup", handler.AdminSignUp)
		r.PATCH("/user/adminverifyemail/:token", handler.AdminVerifyEmail)
		r.POST("/user/kitchenstafflogin", handler.LoginKitchenStaffHandler)
		r.POST("/user/benefactorlogin", handler.LoginFoodBenefactorHandler)
		r.POST("/user/adminlogin", handler.LoginAdminHandler)
		r.POST("/user/beneficiaryforgotpassword", handler.FoodBeneficiaryForgotPassword)
		r.PATCH("/user/beneficiaryresetpassword/:token", handler.FoodBeneficiaryResetPassword)
		r.POST("/user/kitchenstaffforgotpassword", handler.KitchenStaffForgotPassword)
		r.PATCH("/user/kitchenstaffresetpassword/:token", handler.KitchenStaffResetPassword)
		r.POST("/user/adminforgotpassword", handler.AdminForgotPassword)
		r.PATCH("/user/adminresetpassword/:token", handler.AdminResetPassword)
	}

	// authorizeKitchenStaff authorizes all authorized kitchen staff handler
	authorizeKitchenStaff := r.Group("/staff")
	authorizeKitchenStaff.Use(middleware.AuthorizeKitchenStaff(userService.FindKitchenStaffByEmail, userService.TokenInBlacklist))
	{
		authorizeKitchenStaff.POST("/kitchenstafflogout", handler.KitchenStaffLogout)
		authorizeKitchenStaff.PUT("changefoodstatus", handler.ChangeFoodStatus)
	}

	// authorizeBenefactor authorizes all authorized benefactor handler
	authorizeBenefactor := r.Group("/benefactor")
	authorizeBenefactor.Use(middleware.AuthorizeFoodBenefactor(userService.FindFoodBenefactorByEmail, userService.TokenInBlacklist))
	{
		authorizeBenefactor.POST("/beneficiarylogout", handler.FoodBeneficiaryLogout)
		authorizeBenefactor.GET("/brunch", handler.GetBrunchHandle)
		authorizeBenefactor.GET("/dinner", handler.GetDinnerHandle)
		authorizeBenefactor.GET("/getusers", handler.GetUsers)
	}

	// authorizeAdmin authorizes all authorized admin handler
	authorizeAdmin := r.Group("/admin")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(userService.FindAdminByEmail, userService.TokenInBlacklist))
	{
		authorizeAdmin.POST("/createtimetable", handler.CreateFoodTimetableHandle)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
