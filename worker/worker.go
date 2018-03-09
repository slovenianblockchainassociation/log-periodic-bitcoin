package worker

import (
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/regression"
	"os"
	"fmt"
)

const N = 1e5

type Result struct {
	N           int
	J           float64
	Params      *regression.Parameters
	StartParams *regression.Parameters
}

func writeResults(result *Result) error {
	f, err := os.OpenFile("experiments.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%d;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f;%.4f\n", result.N, result.J, result.Params.A, result.Params.B, result.Params.Tc, result.Params.Beta, result.Params.C, result.Params.Omega, result.Params.Phi, result.StartParams.A, result.StartParams.B, result.StartParams.Tc, result.StartParams.Beta, result.StartParams.C, result.StartParams.Omega, result.StartParams.Phi)); err != nil {
		return err
	}
	return nil
}

type Worker interface {
	Start(dataSet []models.DataPoint, resultChn chan *Result)
	Work(dataSet []models.DataPoint) *Result
}
