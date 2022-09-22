package controllers

import (
	"fmt"
	"gotrading/config"
	"gotrading/app/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {

	limit := 100
	duration := "1s"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandole(config.Config.ProductCode, durationTime, limit)
	err := templates.ExecuteTemplate(w, "google.html", df.Candoles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
