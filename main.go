package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TempResponse struct {
	TempCelsius    int32
	TempFahrenheit float32
	Feeling        string
}

func getTemperature() TempResponse {
	//-50 to 55
	temp := rand.Int31n(106) - 50
	var feeling string
	switch {
	case temp <= -30:
		feeling = "Freezing"
	case temp > -30 && temp < -10:
		feeling = "Very cold"
	case temp > -10 && temp < 0:
		feeling = "Cold"
	case temp >= 0 && temp <= 10:
		feeling = "Normal"
	case temp > 10 && temp <= 30:
		feeling = "Warm"
	case temp > 30:
		feeling = "Hot"
	default:
		feeling = "Unknown"
	}

	tempF := covertToFahrenheit(temp)

	resp := TempResponse{TempCelsius: temp, TempFahrenheit: tempF, Feeling: feeling}
	return resp
}

func covertToFahrenheit(temp int32) float32 {
	return (float32(temp) * 1.6) + 32
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		resp := getTemperature()
		c.JSON(http.StatusOK, resp)
	})

	fmt.Println("Started server on http://localhost:8080")
	router.Run(":8080")
}
