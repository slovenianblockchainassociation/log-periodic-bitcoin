package main

import (
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/worker"
	"math"
	"flag"
	"strings"
	"log-periodic-bitcoin/config"
)

func main() {

	inputFile := flag.String("inputFile", config.DefaultDataSetFilePath, "Input dataset file.")
	minDate := flag.Float64("minDate", config.DefaultMinDate, "Dataset will be cut off before this date. Example: 0 = 1.1.2000.")
	maxDate := flag.Float64("maxDate", config.DefaultMaxDate, "Dataset will be cut off after this date. Example: 50 = 1.1.2050.")

	dataSet, err := models.LoadDataSet(*inputFile, *minDate, *maxDate)
	if err != nil {
		panic(err)
	}

	mode := flag.String("mode", config.DefaultSearchMode, "Search mode: full, basic (default), periodic.")
	workers := flag.Int("nWorkers", config.DefaultNumberOfWorkers, "Number of worker processes (default=1).")

	results := make(chan *worker.Result)

	if strings.Compare(*mode, "basic") == 0 {

		for i := 0; i < *workers; i++ {
			clueless := worker.New(results)
			go clueless.StartBasicSearch(dataSet)
		}

	} else if strings.Compare(*mode, "periodic") == 0 {

		A := flag.Float64("A", math.NaN(), "Parameter A of the model. Set if starting periodic search.")
		B := flag.Float64("B", math.NaN(), "Parameter B of the model. Set if starting periodic search.")
		Tc := flag.Float64("Tc", math.NaN(), "Parameter Tc of the model. Set if starting periodic search.")
		Beta := flag.Float64("Beta", math.NaN(), "Parameter Beta of the model. Set if starting periodic search.")

		if math.IsNaN(*A) || math.IsNaN(*B) ||math.IsNaN(*Tc) ||math.IsNaN(*Beta) {
			panic("Set all basic parameters if starting periodic search.")
		}

		for i := 0; i < *workers; i++ {
			clueless := worker.New(results)
			go clueless.StartPeriodicSearch(*A, *B, *Tc, *Beta, dataSet)
		}

	} else if strings.Compare(*mode, "full") == 0 {
		for i := 0; i < *workers; i++ {
			clueless := worker.New(results)
			go clueless.StartFullSearch(dataSet)
		}

	} else {
		panic("Invalid search mode requested.")
	}

	f, err := worker.OpenResultFile(*mode)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	minCost := math.MaxFloat64
	// loop forever, wait for results and write the best in a file
	for {
		select {
		case result := <- results:
			if result.J < minCost {
				minCost = result.J
				err = result.WriteResults(f)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
