package test_functions

import (
	"glowworm_swarm_optimization/gso"
	"math"
)

// x in -5.12 5.12
func GetRastrygin(n int) (gso.Function, float64, float64) {
	return gso.Function(func(x []float64) float64 {
		res := float64(10 * len(x))
		for i := 0; i < len(x); i++ {
			res += math.Pow(x[i], 2) - 10*math.Cos(2*math.Pi*x[i])
		}
		return -res
	}), -5.12, 5.12
}
