package regression

import (
	"log-periodic-bitcoin/models"
	"math"
)

// cos(Omega*ln(Tc - t) + Phi)
func cf(t, tc, omega, phi float64) float64 {
	return math.Cos(omega*math.Log(tc-t) + phi)
}

// sin(Omega*ln(Tc - t) + Phi)
func sf(t, tc, omega, phi float64) float64 {
	return math.Sin(omega*math.Log(tc-t) + phi)
}

// f(t) = A + B*(Tc - t)^Beta * (1 + C*cos(Omega*ln(Tc - t) + Phi))
func f(t float64, p *Parameters) float64 {
	return p.A + p.B*math.Pow(p.Tc-t, p.Beta)*(1+p.C*cf(t, p.Tc, p.Omega, p.Phi))
}

// cost function
// J = 1/2m sum^m (f(t)-y)^2
func J(data []models.DataPoint, p *Parameters) float64 {
	var out float64
	for _, d := range data {
		out += math.Pow(f(d.Date, p)-d.Price, 2)
	}
	return out / float64(len(data)) / 2
}
