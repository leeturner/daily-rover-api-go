package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestInitServerIsNotNil(t *testing.T) {
	httpServer := InitServer("0")
	if httpServer == nil {
		t.Error("Expected httpServer to not be nil")
	}
}

func TestInitServerSetsPort(t *testing.T) {
	port := "0"
	httpServer := InitServer(port)
	if httpServer.port != port {
		t.Errorf("Expected httpServer.port to be %s but got %s", port, httpServer.port)
	}
}

func TestInitServerSetsEchoServer(t *testing.T) {
	httpServer := InitServer("0")
	if httpServer.echoServer == nil {
		t.Error("Expected httpServer.echoServer to not be nil")
	}
}

func TestStatusEndpoint(t *testing.T) {
	// create the server
	httpServer := InitServer("0")

	// create a http client and make a request to the status endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := httpServer.echoServer.NewContext(req, rec)
	c.SetPath("/status")
	err := httpServer.statusEndpoint(c)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d but got %d", http.StatusOK, rec.Code)
	}
	if strings.TrimSpace(rec.Body.String()) != `{"status":200}` {
		t.Errorf("Expected body to be %s but got %s", `{"status":200}`, rec.Body.String())
	}
}

func TestYesterdaysMarsRoverImagesEndpoint(t *testing.T) {
	// create the server
	httpServer := InitServer("0")

	// create a http client and make a request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := httpServer.echoServer.NewContext(req, rec)
	c.SetPath("/v1/photos/")
	err := httpServer.yesterdaysMarsRoverImagesEndpoint(c)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err)
	}

	yesterday := time.Now().Add(-24 * time.Hour).Format(time.DateOnly)
	expectedResponse := "Fetching Mars Rover images for " + yesterday
	if strings.TrimSpace(rec.Body.String()) != expectedResponse {
		t.Errorf("Expected body to be %s but got %s", expectedResponse, rec.Body.String())
	}
}

func TestMarsRoverImagesForEarthDateEndpointFutureDate(t *testing.T) {
	// create the server
	httpServer := InitServer("0")

	// create a date that is in the future
	futureDate := time.Now().Add(24 * time.Hour).Format(time.DateOnly)

	// create a http client and make a request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := httpServer.echoServer.NewContext(req, rec)
	c.SetPath("/v1/photos/:earthDate")
	c.SetParamNames("earthDate")
	c.SetParamValues(futureDate)
	err := httpServer.marsRoverImagesForEarthDateEndpoint(c)
	if err == nil {
		t.Errorf("Expected err not to be nil but got %s", err)
	}

	expectedStatusCode := http.StatusBadRequest
	expectedErrorMessage := fmt.Sprintf("Invalid date - %s - date should be today or in the past", futureDate)
	if err.(*echo.HTTPError).Code != expectedStatusCode {
		t.Errorf("Expected status code to be '%d' but got '%d'", expectedStatusCode, err.(*echo.HTTPError).Code)
	}
	if err.(*echo.HTTPError).Message != expectedErrorMessage {
		t.Errorf("Expected status code to be '%s' but got '%s'", expectedErrorMessage, err.(*echo.HTTPError).Message)
	}
}

func TestMarsRoverImagesForEarthDateEndpointInvalidDate(t *testing.T) {
	// create the server
	httpServer := InitServer("0")

	// create a date that is in the future
	invalidDate := "xxxx-xx-xx"

	// create a http client and make a request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := httpServer.echoServer.NewContext(req, rec)
	c.SetPath("/v1/photos/:earthDate")
	c.SetParamNames("earthDate")
	c.SetParamValues(invalidDate)
	err := httpServer.marsRoverImagesForEarthDateEndpoint(c)
	if err == nil {
		t.Errorf("Expected err not to be nil but got %s", err)
	}

	expectedStatusCode := http.StatusBadRequest
	expectedErrorMessage := fmt.Sprintf("Invalid date - %s - date should be in the format YYYY-MM-DD", invalidDate)
	if err.(*echo.HTTPError).Code != expectedStatusCode {
		t.Errorf("Expected status code to be '%d' but got '%d'", expectedStatusCode, err.(*echo.HTTPError).Code)
	}
	if err.(*echo.HTTPError).Message != expectedErrorMessage {
		t.Errorf("Expected status code to be '%s' but got '%s'", expectedErrorMessage, err.(*echo.HTTPError).Message)
	}
}

func TestMarsRoverImagesForEarthDateEndpoint(t *testing.T) {
	// create the server
	httpServer := InitServer("0")

	testEarthDate := "2021-01-01"

	// create a http client and make a request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := httpServer.echoServer.NewContext(req, rec)
	c.SetPath("/v1/photos/:earthDate")
	c.SetParamNames("earthDate")
	c.SetParamValues(testEarthDate)
	err := httpServer.marsRoverImagesForEarthDateEndpoint(c)
	if err != nil {
		t.Errorf("Expected err to be nil but got %s", err)
	}

	expectedResponse := "Fetching Mars Rover images for " + testEarthDate
	if strings.TrimSpace(rec.Body.String()) != expectedResponse {
		t.Errorf("Expected body to be %s but got %s", expectedResponse, rec.Body.String())
	}
}
