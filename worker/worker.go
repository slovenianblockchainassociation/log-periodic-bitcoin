package worker

import (
	"log-periodic-bitcoin/regression"
	"math"
	"log-periodic-bitcoin/models"
	"os"
	"fmt"
	"log-periodic-bitcoin/config"
	"time"
)

type Result struct {
	N           int64
	J           float64
	Params      *regression.Parameters
	ExeTime     int64
}

func (r *Result) WriteResults(f *os.File) error {
	_, err := f.WriteString(fmt.Sprintf(config.ResultFormat, r.J, r.Params.A, r.Params.B, r.Params.Tc, r.Params.Beta, r.Params.C, r.Params.Omega, r.Params.Phi))
	if err != nil {
		return err
	}
	return nil
}

func OpenResultFile(mode string) (*os.File, error) {
	f, err := os.OpenFile(mode + config.ResultFileSufix, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type Worker struct {
	resultChn chan<- *Result
}

// this is a random search implementation
func New(resultChn chan<- *Result) *Worker {
	return &Worker{resultChn}
}

func (w *Worker) StartBasicSearch(dataSet []models.DataPoint) {
	for {
		result := w.FindBasicParameters(dataSet)
		w.resultChn <- result
	}
}

func (w *Worker) StartPeriodicSearch(a, b, tc, beta float64, dataSet []models.DataPoint) {
	for {
		result := w.FindPeriodicParameters(a, b, tc, beta, dataSet)
		w.resultChn <- result
	}
}

func (w *Worker) StartFullSearch(dataSet []models.DataPoint) {
	for {
		result := w.FindFullParameters(dataSet)
		w.resultChn <- result
	}
}

func (w *Worker) FindBasicParameters(dataSet []models.DataPoint) *Result {
	result := &Result{J: math.MaxFloat64}

	start := time.Now().Unix()
	for result.N < config.NWorkerIterations {
		tmpParams := regression.InitRandomBasicParameters()
		cost := regression.J(dataSet, tmpParams)
		if cost < result.J {
			result.J = cost
			result.Params = tmpParams
		}
		result.N++
	}
	result.ExeTime = time.Now().Unix() - start

	return result
}

func (w *Worker) FindPeriodicParameters(a, b, tc, beta float64, dataSet []models.DataPoint) *Result {
	result := &Result{J: math.MaxFloat64}

	start := time.Now().Unix()
	for result.N < config.NWorkerIterations {
		tmpParams := regression.InitRandomPeriodicParameters(a, b, tc, beta)
		cost := regression.J(dataSet, tmpParams)
		if cost < result.J {
			result.J = cost
			result.Params = tmpParams
		}
		result.N++
	}
	result.ExeTime = time.Now().Unix() - start

	return result
}

func (w *Worker) FindFullParameters(dataSet []models.DataPoint) *Result {
	result := &Result{J: math.MaxFloat64}

	start := time.Now().Unix()
	for result.N < config.NWorkerIterations {
		tmpParams := regression.InitRandomFullParameters()
		cost := regression.J(dataSet, tmpParams)
		if cost < result.J {
			result.J = cost
			result.Params = tmpParams
		}
		result.N++
	}
	result.ExeTime = time.Now().Unix() - start

	return result
}

