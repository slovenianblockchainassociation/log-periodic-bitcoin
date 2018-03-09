package worker

import (
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/regression"
)

const N = 1e5

type Result struct {
	N           int
	J           float64
	Params      *regression.Parameters
}

type Worker interface {
	Start(dataSet []models.DataPoint, resultChn chan *Result)
	Work(dataSet []models.DataPoint) *Result
}
