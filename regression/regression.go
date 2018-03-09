package regression

import (
	"log-periodic-bitcoin/models"
	"math"
)

const Eps = 1e-4
const N = 1e4

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

// derivative of J over A
func dJdA(t, y float64, p *Parameters) float64 {
	return f(t, p) - y
}

// derivative of J over B
func dJdB(t, y float64, p *Parameters) float64 {
	return (f(t, p) - y) * math.Pow(p.Tc-t, p.Beta) * (1 + p.C*cf(t, p.Tc, p.Omega, p.Phi))
}

// derivative of J over Beta
func dJdBeta(t, y float64, p *Parameters) float64 {
	return (f(t, p) - y) * p.B * math.Pow(p.Tc-t, p.Beta) * math.Log(p.Tc-t) * (1 + p.C*cf(t, p.Tc, p.Omega, p.Phi))
}

// derivative of J over Tc
func dJdTc(t, y float64, p *Parameters) float64 {
	return (f(t, p) - y) * p.B * math.Pow(p.Tc-t, p.Beta-1) * (p.Beta*(1+p.C*cf(t, p.Tc, p.Omega, p.Phi)) - p.C*p.Omega*sf(t, p.Tc, p.Omega, p.Phi))
}

// derivative of J over C
func dJdC(t, y float64, p *Parameters) float64 {
	return (f(t, p) - y) * p.B * math.Pow(p.Tc-t, p.Beta) * cf(t, p.Tc, p.Omega, p.Phi)
}

// derivative of J over Omega
func dJdOmega(t, y float64, p *Parameters) float64 {
	return -(f(t, p) - y) * p.B * p.C * p.Omega * math.Pow(p.Tc-t, p.Beta-1) * sf(t, p.Tc, p.Omega, p.Phi)
}

// derivative of J over Phi
func dJdPhi(t, y float64, p *Parameters) float64 {
	return -(f(t, p) - y) * p.B * p.C * math.Pow(p.Tc-t, p.Beta) * sf(t, p.Tc, p.Omega, p.Phi)
}

func Update(data []models.DataPoint, p *Parameters, learningRate *LearningRate, full bool) {
	var da, db, dbeta, dtc, dc, domega, dphi float64
	m := float64(len(data))
	for _, d := range data {
		da += dJdA(d.Date, d.Price, p)
		db += dJdB(d.Date, d.Price, p)
		dbeta += dJdBeta(d.Date, d.Price, p)
		dtc += dJdTc(d.Date, d.Price, p)
		if full {
			dc += dJdC(d.Date, d.Price, p)
			domega += dJdOmega(d.Date, d.Price, p)
			dphi += dJdPhi(d.Date, d.Price, p)
		}
	}
	p.A -= learningRate.A * da / m
	p.B -= learningRate.B * db / m
	p.Beta -= learningRate.Beta * dbeta / m
	p.Tc -= learningRate.Tc * dtc / m
	if full {
		p.C -= learningRate.C * dc /m
		p.Omega -= learningRate.Omega * domega / m
		p.Phi -= learningRate.Phi * dphi / m
	}
	//fmt.Println(da, db, dbeta, dtc, dc, domega, dphi)
}
