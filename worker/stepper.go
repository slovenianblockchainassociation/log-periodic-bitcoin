package worker

import (
	"log-periodic-bitcoin/regression"
	"log-periodic-bitcoin/models"
	"math"
	"fmt"
)

type Stepper struct {
	StepSize *regression.StepSizes
}

// this implementation tries to find a local minimum by changing parameters by step size in the direction that yields lower cost
func NewStepper(stepSizes *regression.StepSizes) *Stepper {
	return &Stepper{stepSizes}
}

func (s *Stepper) Work(dataSet []models.DataPoint) *Result {
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

	var costs []float64

	i := 0
	pcst := math.MaxFloat64
	ccst := 0.0
	for math.Abs(pcst-ccst) > regression.Eps && i < regression.N {
		pcst = ccst
		ccst = regression.J(dataSet, params)
		costs = append(costs, ccst)
		i++
		fmt.Println(ccst, params)
		regression.Step(dataSet, params, s.StepSize, false)
	}

	return &Result{
		len(costs),
		costs[len(costs)-1],
		params,
		startParams,
	}
}