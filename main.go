package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Day struct {
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	q := "Bangalore"
	apiKey := os.Getenv("API_KEY")

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + q)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("weather api not availabe")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	json.Unmarshal(body, &weather)
	location, temp, current, forecast := weather.Location.Name, weather.Current.TempC, weather.Current.Condition.Text, weather.Forecast.Forecastday[0].Day.Condition.Text
	color.Green("City:%s\nTemp:%.1fC \nCurrent:%s\nForecast:%s\n", location, temp, current, forecast)
}
