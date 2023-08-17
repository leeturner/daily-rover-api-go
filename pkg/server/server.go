package server

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const basePhotosApiPath = "/v1/photos"

type Server struct {
	port       string
	echoServer *echo.Echo
}

type status struct {
	Status int `json:"status"`
}

func InitServer(port string) *Server {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())

	s := &Server{
		port:       port,
		echoServer: e,
	}
	s.AddEndpoints()
	return s
}

func (server *Server) statusEndpoint(c echo.Context) error {
	return c.JSON(http.StatusOK, &status{Status: http.StatusOK})
}

func getMarsRoverImagesForDate(earthDate string) string {
	log := "Fetching Mars Rover images for " + earthDate
	fmt.Println(log)
	return log
}

func (server *Server) marsRoverImagesForEarthDateEndpoint(c echo.Context) error {
	earthDate := c.Param("earthDate")
	// Parse the date string using the specified layout
	parsedDate, err := time.Parse(time.DateOnly, earthDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date - "+earthDate+" - date should be in the format YYYY-MM-DD")
	}
	if parsedDate.After(time.Now()) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date - "+earthDate+" - date should be today or in the past")
	}
	return c.String(http.StatusOK, getMarsRoverImagesForDate(earthDate))
}

func (server *Server) yesterdaysMarsRoverImagesEndpoint(c echo.Context) error {
	// Get the current date and time and minus 24 hours to get yesterday's date
	yesterday := time.Now().Add(-24 * time.Hour).Format(time.DateOnly)
	return c.String(http.StatusOK, getMarsRoverImagesForDate(yesterday))
}

func (server *Server) AddEndpoints() {
	server.echoServer.GET(basePhotosApiPath+"/", server.yesterdaysMarsRoverImagesEndpoint)
	server.echoServer.GET(basePhotosApiPath+"/:earthDate", server.marsRoverImagesForEarthDateEndpoint)
	server.echoServer.GET("/status", server.statusEndpoint)
}

func (server *Server) Start() error {
	return server.echoServer.Start(fmt.Sprintf(":%s", server.port))
}
