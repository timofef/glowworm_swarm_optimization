package test_functions

import (
	"glowworm_swarm_optimization/utils"
	"math"
)

func GetRastrygin(n int) (Function, utils.Interval, []float64, float64) {
	return Function(func(x []float64) float64 {
		res := float64(10 * len(x))
		for i := 0; i < len(x); i++ {
			res += x[i]*x[i] - 10*math.Cos(2*math.Pi*x[i])
		}
		return -res
	}), utils.Interval{Min: -5.0, Max: 5.0}, make([]float64, n, n), 0.0
}
