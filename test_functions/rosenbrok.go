package test_functions

import (
	"glowworm_swarm_optimization/gso"
	"math"
)

//
func GetRosenbrok(n int) (gso.Function, float64, float64) {
	return gso.Function(func(x []float64) float64 {
		var res float64
		for i := 0; i < len(x)-1; i++ {
			res += 100*(x[i+1]-math.Pow(x[i], 2)) + math.Pow(x[i]-1, 2)
		}
		return res
	}), -5.0, 5.0
}
