package routes

import (
	"fmt"
	job "go-api/app/job/handler"
	JobInterface "go-api/app/job/service"
	user "go-api/app/user/handler"
	UserInterface "go-api/app/user/service"
	"go-api/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	authorizeMiddleware echo.MiddlewareFunc
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &Validator{validator: validator.New()}
	e.HTTPErrorHandler = HttpErrorHandler

	conf := config.GetConfig()
	authorizeMiddleware = middleware.JWT([]byte(conf.GetString("jwt.secret_key")))

	return e
}

func HttpErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
			}

			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}

func NewUserHandler(router *echo.Echo, uc UserInterface.IUserService) {
	h := &user.UserHandler{
		Service: uc,
	}

	user := router.Group("/user")
	{
		user.POST("/login", h.Login)
	}
}

func NewJobHandler(router *echo.Echo, ps JobInterface.IJobService) {
	h := &job.JobHandler{
		Service: ps,
	}

	job := router.Group("/job", authorizeMiddleware)
	{
		job.GET("/list", h.ListJob)
		job.GET("/list/:id", h.ListJobByID)
	}
}
