package worker

import (
	"log-periodic-bitcoin/regression"
	"fmt"
	"math"
	"log-periodic-bitcoin/models"
)

type Regressor struct {
	LearningRate *regression.LearningRate
}

// gradiend descent implementation
func NewRegressor(learningRate *regression.LearningRate) *Regressor {
	return &Regressor{learningRate}
}

func (r *Regressor) Work(dataSet []models.DataPoint) *Result {
	params := regression.InitRandomParameters(false)
	startParams := &regression.Parameters{
		A:     params.A,
		B:     params.B,
		Tc:    params.Tc,
		Beta:  params.Beta,
		C:     params.C,
		Omega: params.Omega,
		Phi:   params.Phi,
	}
	fmt.Println(params)
	var costs []float64

	i := 0
	pcst := math.MaxFloat64
	ccst := 0.0
	for math.Abs(pcst-ccst) > regression.Eps && i < regression.N {
		pcst = ccst
		ccst = regression.J(dataSet, params)
		costs = append(costs, ccst)
		i++
		//fmt.Println(ccst, params)
		regression.Update(dataSet, params, r.LearningRate, false)
	}

	return &Result{
		len(costs),
		costs[len(costs)-1],
		params,
		startParams,
	}
}
