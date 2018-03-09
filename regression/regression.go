package regression

import (
	"log-periodic-bitcoin/models"
	"math"
)

// f(t) = A + B*(Tc - t)^Beta
func smallf(t float64, A, B, Tc, Beta float64) float64 {
	return A + B*math.Pow(Tc-t, Beta)
}

// cost function
// J = 1/2m sum^m (f(t)-y)^2
func SmallJ(A, B, Tc, Beta float64, data []models.DataPoint) float64 {
	var out float64
	for _, d := range data {
		out += math.Pow(smallf(d.Date, A, B, Tc, Beta)-d.Price, 2)
	}
	return out / float64(len(data)) / 2
}

// f(t) = A + B*(Tc - t)^Beta * (1 + C*cos(Omega*ln(Tc - t) + Phi))
func f(t float64, A, B, Tc, Beta, C, Omega, Phi float64) float64 {
	return A + B*math.Pow(Tc-t, Beta)*(1+C*math.Cos(Omega*math.Log(Tc-t) + Phi))
}

// cost function
// J = 1/2m sum^m (f(t)-y)^2
func J(A, B, Tc, Beta, C, Omega, Phi float64, data []models.DataPoint) float64 {
	var out float64
	for _, d := range data {
		out += math.Pow(f(d.Date, A, B, Tc, Beta, C, Omega, Phi)-d.Price, 2)
	}
	return out / float64(len(data)) / 2
}
