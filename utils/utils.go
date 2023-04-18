package utils

import "math"

type OptimizationConfig struct {
	MaxTime int      // Time steps limit
	DeltaF  float64  // Corridor width
	DeltaT  int      // Corridor length
	N       int      // Swarm size
	M       int      // Dimension of solution
	Diap    Interval // Diap for x vector
}

type Interval struct {
	Min float64
	Max float64
}

func Norm(a, b []float64) float64 {
	var sum float64
	for i := range a {
		c := a[i] - b[i]
		sum += c * c
	}
	return math.Sqrt(sum)
}
