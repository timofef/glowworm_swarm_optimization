package test_functions

import (
	"glowworm_swarm_optimization/utils"
	"math"
)

func GetRosenbrok(n int) (Function, utils.Interval, []float64, float64) {
	v := make([]float64, n, n)
	for i := 0; i < n; i++ {
		v[i] = 1.0
	}
	return Function(func(x []float64) float64 {
		var res float64
		for i := 0; i < len(x)-1; i++ {
			res += 100*math.Pow(x[i+1]-math.Pow(x[i], 2), 2) + math.Pow(1.0-x[i], 2)
		}
		return -res
	}), utils.Interval{Min: -100.0, Max: 100.0}, v, 0.0
}
