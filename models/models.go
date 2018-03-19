package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

	return float64(t.Year()) + (float64(t.YearDay()) + float64(t.Hour()) / 24)/float64(year.YearDay()) - 2000
}

func limitDataSetByDate(minDate, maxDate float64, dataSet []DataPoint) ([]DataPoint, error) {
	if maxDate-minDate <= 0 {
		return nil, errors.New("maxDate before minDate")
	}
	if minDate > dataSet[len(dataSet)-1].Date {
		return nil, errors.New("minDate after last point in dataSet")
	}
	if maxDate < dataSet[0].Date {
		return nil, errors.New("maxDate before first point in dataSet")
	}
	start := -1
	for i, v := range dataSet {
		if start < 0 && v.Date > minDate {
			start = i
		}
		if v.Date > maxDate {
			return dataSet[start:i], nil
		}
	}
	return dataSet[start:], nil
}

func LoadDataSet(filename string, minDate, maxDate float64) ([]DataPoint, error) {
	// read data file
	rawData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// unmarshal rawData to data points
	var dataSet []DataPoint
	err = json.Unmarshal(rawData, &dataSet)
	if err != nil {
		return nil, err
	}

	// limit dataSet by minDate and maxDate
	dataSet, err = limitDataSetByDate(minDate, maxDate, dataSet)
	if err != nil {
		return nil, err
	}
	return dataSet, nil
}
