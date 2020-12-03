package Routes

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Controllers"
	"theapp/Middlewares"
	"theapp/Service"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()
	r.Use(Middlewares.SignRequest)
	r.Use(sentrygin.New(sentrygin.Options{}))

	var loginService Service.LoginService = Service.StaticLoginService()
	var jwtService Service.JWTService = Service.JwtAuthService()
	var loginController Controllers.LoginController = Controllers.LoginHandler(loginService, jwtService)

	r.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != ""{
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}else{
			context.JSON(http.StatusUnauthorized, nil)
		}
	})

	usergroup := r.Group("/userapi")
	{
		usergroup.GET("user", Controllers.GetUsers)

	}

	admingroup := r.Group("/admin",  Middlewares.AuthorizeJWT(), Middlewares.AuthorizeUser())
	{
		admingroup.GET("user/:id", Controllers.GetUserByID)
		admingroup.POST("user", Controllers.CreateUser)
		admingroup.PUT("user/:id", Controllers.UpdateUser)
		admingroup.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}