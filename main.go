package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(getLogger())
	e.Renderer = Renderer()
	e.Static("/static", "public/static/")
	e.GET("/api", apiHandler).Name = "api"
	e.GET("/", indexHandlerFactory()).Name = "index"
	e.Logger.Fatal(e.Start(":8000"))
}

func getLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}

func apiHandler(c echo.Context) error {

	event, err := FromEchoContext(c)

	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	c.Response().Header().Add("Content-Disposition", getResponseFileName(event.Title, event.Start))
	c.Response().Header().Set("Content-Type", "text/calendar")
	return c.Render(http.StatusOK, "event.ics", event)
}

func getResponseFileName(title string, start time.Time) string {
	return fmt.Sprintf("attachment; filename=\"%s-%s.ics\"", title, start.Format("20060102T0303pm"))
}

func indexHandlerFactory() func(c echo.Context) error {
	timezone, err := readTimezoneFromFS()
	list := timezone.GetTimeList()
	if err != nil {
		return nil
	}
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"zones": list,
		})
	}
}
