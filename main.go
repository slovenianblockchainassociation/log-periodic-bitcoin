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

func limitDataSetByMaxDate(maxDate float64, dataSet []models.DataPoint) []models.DataPoint {
	for i, v := range dataSet {
		if v.Date > maxDate {
			return dataSet[:i]
		}
	}
	return dataSet
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

	dataSet = limitDataSetByMaxDate(17.95, dataSet)

	// random search code starts here
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
		clueless := worker.NewClueless()
		go clueless.Start(dataSet, results)
	}

	f, err := os.OpenFile("randomSearch.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		select {
		case result := <- results:
			if result.J >= minCost {
				continue
			}
			minCost = result.J
			if _, err = f.WriteString(fmt.Sprintf("%.4f;%.2f %.2f %.2f %.2f\n", result.J, result.Params.A, result.Params.B, result.Params.Tc, result.Params.Beta)); err != nil {
				panic(err)
			}
		}
	}
	// random search code ends here

	//learningRate := &regression.LearningRate{
	//	A: 10e-7,
	//	B: 10e-7,
	//	Tc: 10e-7,
	//	Beta: 10e-7,
	//}

	//stepSizes := &regression.StepSizes{
	//	A: 10,
	//	B: 10,
	//	Tc: 0.01,
	//	Beta: 0.01,
	//}

	//regressor := worker.NewRegressor(learningRate)
	//stepper := worker.NewStepper(stepSizes)
	//clueless := worker.NewClueless()

	//result := regressor.Work(dataSet)
	//result := stepper.Work(dataSet)
	//result := clueless.Work(dataSet)

	//if math.IsNaN(result.J) {
	//	fmt.Println("Try diverged")
	//	return
	//}
	//
	//err = worker.writeResults(result)
	//if err != nil {
	//	panic(err)
	//}
}
