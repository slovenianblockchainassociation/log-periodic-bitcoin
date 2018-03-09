package main

import (
	"encoding/json"
	"io/ioutil"
	"log-periodic-bitcoin/models"
	"fmt"
	"log-periodic-bitcoin/worker"
	"math"
	"os"
	"strconv"
)

var paramRanges = map[string][3]float64{
	// min, max, step, portion (calculated based on number of workers)
	"A": {0, 100, 0.1},
	"B": {-100, 0, 0.1},
	"Tc": {18, 18.3, 0.01},
	"Beta": {0.05, 0.25, 0.01},
	"C": {0, 0.1, 0.01},
	"Omega": {10, 30, 0.1},
	"Phi": {0, 2*math.Pi, 0.1},
}

func main() {

	// read data file
	rawData, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	// unmarshal rawData to data points
	var dataSet []models.DataPoint
	err = json.Unmarshal(rawData, &dataSet)
	if err != nil {
		panic(err)
	}

	dataSet = models.LimitDataSetByMaxDate(17.95, dataSet)

	results := make(chan *worker.Result)
	minCost := math.MaxFloat64

	workers := 1

	if len(os.Args) > 1 {
		workers, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < workers; i++ {
		w := worker.New(i, paramRanges)
		go w.Work(dataSet, results)
	}

	f, err := os.OpenFile("randomSearch.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		select {
		case result := <- results:
			if result.End {
				workers--
				if workers == 0 {
					return
				}
			}
			if result.J >= minCost {
				continue
			}
			minCost = result.J
			if _, err = f.WriteString(fmt.Sprintf("%.4f;%.1f %.1f %.2f %.2f %.2f %.1f %.2f\n", result.J, result.A, result.B, result.Tc, result.Beta, result.C, result.Omega, result.Phi)); err != nil {
				panic(err)
			}
		}
	}
}
