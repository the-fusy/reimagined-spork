package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type DayData struct {
	Temp float32
}

func GetForecast(city string) (map[time.Time]DayData, error) {
	url := "https://api.weatherapi.com/v1/forecast.json?aqu=no&days=10&key=c8a4faae610d4e5bb9404401222701&q=" + city

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Forecast struct {
			ForecastDay []struct {
				Date string `json:"date"`
				Day  struct {
					Temp float32 `json:"avgtemp_c"`
				} `json:"day"`
			} `json:"forecastday"`
		} `json:"forecast"`
	}
	if err = json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	weatherByDay := make(map[time.Time]DayData)
	for _, day := range apiResp.Forecast.ForecastDay {
		date, err := time.Parse("2006-01-02", day.Date)
		if err != nil {
			return nil, err
		}
		weatherByDay[date] = DayData{Temp: day.Day.Temp}
	}

	return weatherByDay, nil
}
