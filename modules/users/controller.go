package users

import (
	"net/http"
	users "online-shop-home-test/modules/users/dto"
	"online-shop-home-test/modules/utils"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	authRoute := rg.Group("/auth")

	authRoute.POST("/register", registerUser)
	authRoute.POST("/login", loginUser)

	authProtectedRoute := authRoute.Use(utils.JwtAuthMiddleware())
	authProtectedRoute.GET("/authorize", authorizeToken)
}

// Auth      		godoc
// @Summary     	Register user
// @Description  	Regist a new user
// @Param			register body users.CreateUserDTO true  "Register user"
// @Tags         	auth
// @Produce      	json
// @Success      	200  {object} users.SuccessRegisterResponse
// @Failure      	400  {object} users.BadRequestRegisterResponse
// @Router       	/auth/register [post]
func registerUser(ctx *gin.Context) {
	var input users.CreateUserDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

// Auth      		godoc
// @Summary     	Login user
// @Description  	Login a new user
// @Param			login body users.LoginDTO true  "Login user"
// @Tags         	auth
// @Produce      	json
// @Success      	200  {object} users.SuccessLoginResponse
// @Failure      	400  {object} users.BadRequestLoginResponse
// @Router       	/auth/login [post]
func loginUser(ctx *gin.Context) {
	var input users.LoginDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := Login(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

// Auth      		godoc
// @Summary     	Authorize user
// @Description  	Authorize user access token after login
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags         	auth
// @Produce      	json
// @Success      	200  {object} users.SuccessAuthorizeToken
// @Failure      	401  {object} users.UnauthorizedToken
// @Router       	/auth/authorize [get]
func authorizeToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "token is valid",
	})
}
