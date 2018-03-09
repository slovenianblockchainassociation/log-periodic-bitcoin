package worker

import (
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/regression"
	"math"
)

type Clueless struct {}

// this is a random search implementation
func NewClueless() *Clueless {
	return &Clueless{}
}

func (c *Clueless) Start(dataSet []models.DataPoint, resultChn chan<- *Result) {
	minCost := math.MaxFloat64
	for {
		result := c.Work(dataSet)
		if result.J < minCost {
			minCost = result.J
			resultChn <- result
		}
	}
}

func (c *Clueless) Work(dataSet []models.DataPoint) *Result {
	params := regression.InitRandomParameters(false)

	minCost := math.MaxFloat64

	i := 0
	for i < N {
		tmpParams := regression.InitRandomParameters(false)
		cost := regression.J(dataSet, tmpParams)
		if cost < minCost {
			minCost = cost
			params = tmpParams
		}
		i++
	}

	return &Result{
		i,
		minCost,
		params,
	}
}
