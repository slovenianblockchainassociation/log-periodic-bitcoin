package models

import (
	"encoding/json"
	"math"
	"time"
)

type DataPoint struct {
	Date  float64
	Price float64
}

func (dp *DataPoint) UnmarshalJSON(b []byte) error {
	var raw struct {
		Date  int64   `json:"date"`
		Price float64 `json:"close"`
	}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}
	dp.Date = UnixToDecimal(raw.Date)
	dp.Price = math.Log(raw.Price)
	return nil
}

// function that converts unix timestamp to decimal time.
// Note that it only takes years and days into account.
// If timeframe lower then daily is required, hourly and lower contributions need to be added
func UnixToDecimal(u int64) float64 {
	t := time.Unix(u, 0)
	year := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.Local)

	return float64(t.Year()) + float64(t.YearDay())/float64(year.YearDay()) - 2000
}
