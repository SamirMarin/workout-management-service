package main

import (
	"github.com/SamirMarin/workout-management-service/internal/workout"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/create", create)
	e.Logger.Fatal(e.Start(":1323"))
}

func create(c echo.Context) error {
	workout := workout.Workout{}
	if err := c.Bind(&workout); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := workout.CreateWorkout()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Workout created")
}
