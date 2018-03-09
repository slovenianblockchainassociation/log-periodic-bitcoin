package regression

import (
	"math"
	"math/big"
	"crypto/rand"
)

// generate random from /dev/urandom
func RandFloat64(r int64, d float64) float64 {
	// r is upper bound of int produced by rand.Int [0,r]
	// d is the end divisor
	nBig, err := rand.Int(rand.Reader, big.NewInt(r))
	if err != nil {
		panic(err)
	}
	return float64(nBig.Int64()) / d
}

// parameters for ln(p(t)) = f(t) = A + B*(Tc - t)^Beta * (1 + C*cos(Omega*ln(Tc - t) + Phi))
type Parameters struct {
	A     float64
	B     float64
	Tc    float64
	Beta  float64
	C     float64
	Omega float64
	Phi   float64
}

func InitParameters(a, b, tc, beta, c, omega, phi float64) *Parameters {
	return &Parameters{
		A:     a,
		B:     b,
		Tc:    tc,
		Beta:  beta,
		C:     c,
		Omega: omega,
		Phi:   phi,
	}
}

func InitRandomParameters(full bool) *Parameters {
	p := &Parameters{
		A:    RandFloat64(10000, 100),
		B:    -RandFloat64(10000, 100),
		Tc:   RandFloat64(100, 50) + 18,
		Beta: RandFloat64(50, 100) + 0.05,
	}
	if full {
		p.C = RandFloat64(100, 100)
		p.Omega = RandFloat64(100, 10)
		p.Phi = RandFloat64(1000, 500 / math.Pi)
	}
	return p
}

type LearningRate struct {
	A     float64
	B     float64
	Tc    float64
	Beta  float64
	C     float64
	Omega float64
	Phi   float64
}

type StepSizes struct {
	A     float64
	B     float64
	Tc    float64
	Beta  float64
	C     float64
	Omega float64
	Phi   float64
}
