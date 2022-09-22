package models

import "time"

type DataFrameCandole struct {
	Productcode string    `json:"product_code"`
	Duration    time.Duration    `json:"duration"`
	Candoles    []Candole `json:"candoles"`
}

func (df *DataFrameCandole) Times() []time.Time {
	s := make([]time.Time, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.Time
	}
	return s
}

func (df *DataFrameCandole) Opens() []float64 {
	s := make([]float64, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.Open
	}
	return s
}
func (df *DataFrameCandole) Closes() []float64 {
	s := make([]float64, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.Close
	}
	return s
}
func (df *DataFrameCandole) Highs() []float64 {
	s := make([]float64, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.High
	}
	return s
}
func (df *DataFrameCandole) Lows() []float64 {
	s := make([]float64, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.Low
	}
	return s
}
func (df *DataFrameCandole) Volumes() []float64 {
	s := make([]float64, len(df.Candoles))
	for i, candole := range df.Candoles {
		s[i] = candole.Volume
	}
	return s
}
