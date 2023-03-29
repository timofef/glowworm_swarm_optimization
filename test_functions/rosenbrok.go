package test_functions

import (
	"glowworm_swarm_optimization/gso"
	"math"
)

func GetRosenbrok(n int) (gso.Function, gso.Interval, []float64, float64) {
	v := make([]float64, n, n)
	for i := 0; i < n; i++ {
		v[i] = 1.0
	}
	return gso.Function(func(x []float64) float64 {
		var res float64
		for i := 0; i < len(x)-1; i++ {
			res += 100*math.Pow(x[i+1]-math.Pow(x[i], 2), 2) + math.Pow(x[i]-1, 2)
		}
		return -res
	}), gso.Interval{Min: -5.0, Max: 5.0}, v, 0.0
}
