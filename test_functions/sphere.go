package test_functions

import (
	"glowworm_swarm_optimization/gso"
	"math"
)

// -5 5
func GetSphere(n int) (gso.Function, float64, float64) {
	return gso.Function(func(x []float64) float64 {
		var res float64
		for i := 0; i < len(x); i++ {
			res += math.Pow(x[i], 2)
		}
		return -res
	}), -5.0, 5.0
}