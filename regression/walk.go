package regression

import (
	"log-periodic-bitcoin/models"
	"math"
)

func Step(data []models.DataPoint, p *Parameters, stepSizes *StepSizes, full bool) {
	var posA, negA float64
	var posB, negB float64
	var posTc, negTc float64
	var posBeta, negBeta float64

	for _, d := range data {
		p.A += stepSizes.A
		posA += math.Pow(f(d.Date, p) - d.Price, 2)
		p.A -= 2*stepSizes.A
		negA += math.Pow(f(d.Date, p) - d.Price, 2)
		p.A += stepSizes.A

		p.B += stepSizes.B
		posB += math.Pow(f(d.Date, p) - d.Price, 2)
		p.B -= 2*stepSizes.B
		negB += math.Pow(f(d.Date, p) - d.Price, 2)
		p.B += stepSizes.B

		p.Tc += stepSizes.Tc
		posTc += math.Pow(f(d.Date, p) - d.Price, 2)
		p.Tc -= 2*stepSizes.Tc
		negTc += math.Pow(f(d.Date, p) - d.Price, 2)
		p.Tc += stepSizes.Tc

		p.Beta += stepSizes.Beta
		posBeta += math.Pow(f(d.Date, p) - d.Price, 2)
		p.Beta -= 2*stepSizes.Beta
		negBeta += math.Pow(f(d.Date, p) - d.Price, 2)
		p.Beta += stepSizes.Beta
	}

	if posA < negA {
		p.A += stepSizes.A
	} else {
		p.A -= stepSizes.A
	}

	if posB < negB {
		p.B += stepSizes.B
	} else {
		p.B -= stepSizes.B
	}

	if posTc < negTc {
		p.Tc += stepSizes.Tc
	} else {
		p.Tc -= stepSizes.Tc
	}

	if posBeta < negBeta {
		p.Beta += stepSizes.Beta
	} else {
		p.Beta -= stepSizes.Beta
	}
}
