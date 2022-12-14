package controllers

import (
	"encoding/json"
	"fmt"
	"gotrading/app/models"
	"gotrading/config"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseFiles("app/views/chart.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "chart.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not Found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		APIError(w, "No product_code", http.StatusBadRequest)
		return
	}
	strlimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strlimit)
	if strlimit == "" || err != nil || limit < 0 || limit > 1000{
		limit = 1000
	}
	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.Config.Durations[duration]

	df, _ := models.GetAllCandle(productCode, durationTime, limit)
	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	http.HandleFunc("/api/candle/", apiMakeHandler(apiCandleHandler))
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
