package worker

import (
	"log-periodic-bitcoin/models"
	"log-periodic-bitcoin/regression"
	"math"
)

const N = 1e5

type Result struct {
	J                             float64
	A, B, Tc, Beta, C, Omega, Phi float64
	End                           bool
}

type Worker struct {
	I     float64
	A     [3]float64
	B     [3]float64
	Tc    [3]float64
	Beta  [3]float64
	C     [3]float64
	Omega [3]float64
	Phi   [3]float64
}

func New(i int, ranges map[string][3]float64) *Worker {
	return &Worker{
		float64(i),
		ranges["A"],
		ranges["B"],
		ranges["Tc"],
		ranges["Beta"],
		ranges["C"],
		ranges["Omega"],
		ranges["Phi"],
	}
}

func (w *Worker) Work(data []models.DataPoint, resultChn chan<- *Result) {
	result := &Result{
		J:   math.MaxFloat64,
		End: false,
	}
	for a := w.A[0] + (w.A[1]-w.A[0])*w.I; a < w.A[0]+(w.A[1]-w.A[0])*(w.I+1); a += w.A[2] {
		for b := w.B[0]; b < w.B[1]; a += w.B[2] {
			for tc := w.Tc[0]; tc < w.Tc[1]; tc += w.Tc[2] {
				for beta := w.Beta[0]; beta < w.Beta[1]; beta += w.Beta[2] {
					for c := w.C[0]; c < w.C[1]; c += w.C[2] {
						for omega := w.Omega[0]; omega < w.Omega[1]; omega += w.Omega[2] {
							for phi := w.Phi[0]; phi < w.Phi[1]; phi += w.Phi[2] {
								cost := regression.J(a, b, tc, beta, c, omega, phi, data)
								if cost < result.J {
									result.J = cost

									result.A = a
									result.B = b
									result.Tc = tc
									result.Beta = beta
									result.C = c
									result.Omega = omega
									result.Phi = phi

									resultChn <- result
								}
							}
						}
					}
				}
			}
		}
	}
	resultChn <- &Result{End: true}
}
