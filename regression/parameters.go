package regression

import (
	"crypto/rand"
	"math"
	"math/big"
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

func InitRandomBasicParameters() *Parameters {
	return &Parameters{
		A:    RandFloat64(10000, 100),
		B:    -RandFloat64(10000, 100),
		Tc:   RandFloat64(100, 50) + 18,
		Beta: RandFloat64(50, 100) + 0.05,
	}
}

func InitRandomPeriodicParameters(a, b, tc, beta float64) *Parameters {
	return &Parameters{
		A:     a,
		B:     b,
		Tc:    tc,
		Beta:  beta,
		C:     RandFloat64(10, 1000) + 0.001,
		Omega: RandFloat64(100, 4) + 5,
		Phi:   RandFloat64(1000, 500/math.Pi),
	}
}

func InitRandomFullParameters() *Parameters {
	return &Parameters{
		A:     RandFloat64(10000, 100),
		B:     -RandFloat64(10000, 100),
		Tc:    RandFloat64(100, 50) + 18,
		Beta:  RandFloat64(50, 100) + 0.05,
		C:     RandFloat64(10, 1000) + 0.01,
		Omega: RandFloat64(100, 4) + 5,
		Phi:   RandFloat64(1000, 500/math.Pi),
	}
}
