package main

import (
	"glowworm_swarm_optimization/gso"
)

func main() {
	conf := gso.OptimizationConfig{
		MaxTime: 1000,
		S:       2,
		Diap:    gso.Interval{Min: -100, Max: 100},
	}

	gso.Optimize(conf, gso.Function{
		N: 2,
		F: func(i int, f float64) float64 {
			return 0.0
		}},
	)
}