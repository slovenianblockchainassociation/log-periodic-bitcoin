package main

import (
	"flag"
	"log"
	"log-periodic-bitcoin/config"
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/worker"
	"math"
	"strings"
)

func main() {

	// flags ---------------------------------------
	inputFile := flag.String("inputFile", config.DefaultDataSetFilePath, "Input dataset file.")
	minDate := flag.Float64("minDate", config.DefaultMinDate, "Dataset will be cut off before this date. Example: 1 = 1.1.2001.")
	maxDate := flag.Float64("maxDate", config.DefaultMaxDate, "Dataset will be cut off after this date. Example: 50 = 1.1.2050.")

	mode := flag.String("mode", config.DefaultSearchMode, "Search mode: full, basic , periodic.")
	workers := flag.Int("workers", config.DefaultNumberOfWorkers, "Number of worker processes.")
	nIter := flag.Int("nIter", config.DefaultNIterations, "Number of iterations per worker.")

	A := flag.Float64("A", math.NaN(), "Parameter A of the model. Set if starting periodic search.")
	B := flag.Float64("B", math.NaN(), "Parameter B of the model. Set if starting periodic search.")
	Tc := flag.Float64("Tc", math.NaN(), "Parameter Tc of the model. Set if starting periodic search.")
	Beta := flag.Float64("Beta", math.NaN(), "Parameter Beta of the model. Set if starting periodic search.")

	reportSpeed := flag.Bool("s", false, "Use this flag to turn on execution speed reporting.")

	flag.Parse()

	// reset maxDate if Tc is set before maxDate
	if !math.IsNaN(*Tc) && *Tc < *maxDate {
		*maxDate = *Tc
	}
	// ----------------------------------------------

	// load data set --------------------------------
	dataSet, err := models.LoadDataSet(*inputFile, *minDate, *maxDate)
	if err != nil {
		panic(err)
	}
	// ----------------------------------------------

	// start workers --------------------------------
	results := make(chan *worker.Result)

	if strings.Compare(*mode, "basic") == 0 {

		for i := 0; i < *workers; i++ {
			clueless := worker.New(int64(*nIter), results)
			go clueless.StartBasicSearch(dataSet)
		}

	} else if strings.Compare(*mode, "periodic") == 0 {

		if math.IsNaN(*A) || math.IsNaN(*B) || math.IsNaN(*Tc) || math.IsNaN(*Beta) {
			panic("Set all basic parameters if starting periodic search.")
		}

		for i := 0; i < *workers; i++ {
			clueless := worker.New(int64(*nIter), results)
			go clueless.StartPeriodicSearch(*A, *B, *Tc, *Beta, dataSet)
		}

	} else if strings.Compare(*mode, "full") == 0 {
		for i := 0; i < *workers; i++ {
			clueless := worker.New(int64(*nIter), results)
			go clueless.StartFullSearch(dataSet)
		}

	} else {
		panic("Invalid search mode requested.")
	}
	// -----------------------------------------------

	// open result file ------------------------------
	f, err := worker.OpenResultFile(*mode, *inputFile, *minDate, *maxDate)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// -----------------------------------------------

	// loop forever, wait for results, write the best to file, report execution speed
	minCost := math.MaxFloat64
	var cumTime int64
	var iterations int64

	for {
		select {
		case result := <-results:
			if result.J < minCost {
				minCost = result.J
				err = result.WriteResults(f)
				if err != nil {
					panic(err)
				}
			}
			if *reportSpeed {
				cumTime += result.ExeTime
				iterations += result.N
				log.Printf("%d seconds for %.0e iterations. Average speed: %d ops/sec.", result.ExeTime, float64(result.N), int64(*workers)*iterations/cumTime)
			}
		}
	}
	// -----------------------------------------------
}
