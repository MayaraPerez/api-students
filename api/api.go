package api

import (
	"github.com/MayaraPerez/api-students/db"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database:= db.Init()
	StudentDb := db.NewStudentHandle(database)

	return &API{
		Echo: e,
		DB: StudentDb,
	}
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) Routes() {
	// Routes
	api.Echo.GET("/students", api.GetStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

