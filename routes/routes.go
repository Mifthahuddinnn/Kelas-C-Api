package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mahasiswa/controllers"
)

func Routes(e *echo.Echo) {
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/students", controllers.GetStudents)
	e.POST("/student", controllers.AddStudent)
	e.GET("/student/:id", controllers.GetStudentByID)
	e.PUT("/student/:id", controllers.UpdateStudent)
	e.DELETE("/student/:id", controllers.DeleteStudent)
}
