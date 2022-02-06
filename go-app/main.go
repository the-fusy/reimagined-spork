package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spadiff/reimagined-spork/internal/weather"
	"strconv"
	"time"
)

func main() {
	app.Route("/", new(MainPage))
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("./docs/go-app", &app.Handler{
		Name: "Test",
		Styles: []string{
			"https://www.w3schools.com/w3css/4/w3.css",
		},
		Resources: app.GitHubPages("reimagined-spork/go-app"),
	})
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
}

type MainPage struct {
	app.Compo
	weather Weather
}

func (m *MainPage) Render() app.UI {
	return app.Div().Class("w3-container").Body(&m.weather)
}

type Weather struct {
	app.Compo

	city             string
	date             time.Time
	minDate, maxDate time.Time

	dayData map[time.Time]weather.DayData
}

func (w *Weather) OnMount(ctx app.Context) {
	w.city = "London"
	w.date = time.Now().UTC().Truncate(24 * time.Hour)
	w.minDate = w.date
	w.maxDate = w.date.Add(2 * 24 * time.Hour)

	var err error
	w.dayData, err = weather.GetForecast(w.city)
	if err != nil {
		w.dayData = nil
	}
}

func (w *Weather) Render() app.UI {
	var temp string
	if w.dayData == nil {
		temp = "ERR"
	} else {
		temp = strconv.FormatFloat(float64(w.dayData[w.date].Temp), 'f', -1, 32)
	}

	weatherBlock := app.Div().Class("w3-col s8").Body(
		app.Div().Class("w3-container w3-row w3-blue").Body(
			app.H1().Class("w3-col s6").Text(w.city),
			app.H1().Class("w3-col s6 w3-right-align").Text(w.date.Format("02 January")),
		),
		app.Div().Class("w3-container").Body(
			app.Div().Class("w3-center").Style("font-size", "200px").Text(temp),
		),
	)

	return app.Div().Class("w3-card-4 w3-display-middle").Class("w3-row").Style("width", "50%").Style("height", "400px").Body(
		app.Div().Class("w3-button").Class("w3-col s2").Style("height", "100%").Text("<").OnClick(w.loadPrevDay),
		weatherBlock,
		app.Div().Class("w3-button").Class("w3-col s2").Style("height", "100%").Text(">").OnClick(w.loadNextDay),
	)
}

func (w *Weather) loadPrevDay(ctx app.Context, e app.Event) {
	w.date = w.date.Add(-24 * time.Hour)
	if w.date.Before(w.minDate) {
		w.date = w.minDate
	}
}

func (w *Weather) loadNextDay(ctx app.Context, e app.Event) {
	w.date = w.date.Add(24 * time.Hour)
	if w.date.After(w.maxDate) {
		w.date = w.maxDate
	}
}
