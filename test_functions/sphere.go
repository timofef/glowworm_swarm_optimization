package test_functions

import (
	"glowworm_swarm_optimization/utils"
)

func GetSphere(n int) (Function, utils.Interval, []float64, float64) {
	return Function(func(x []float64) float64 {
		var res float64
		for i := 0; i < len(x); i++ {
			res += x[i] * x[i]
		}
		return -res
	}), utils.Interval{Min: -5.0, Max: 5.0}, make([]float64, n, n), 0.0
}
