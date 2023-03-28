package main

import (
	"fmt"
	"glowworm_swarm_optimization/gso"
	"glowworm_swarm_optimization/test_functions"
	"math/rand"
)

func multistart(conf gso.OptimizationConfig, f gso.Function, num int) {
	min := -1000000.0

	for i := 0; i < num; i++ {
		_, val := gso.Optimize(conf, f)
		if val > min {
			min = val
		}
	}
	fmt.Printf("Min: %f\n", -min)
}

func main() {
	rand.Seed(69420)

	// Get test function and diap for x vector
	f, minX, maxX := test_functions.GetSphere(2)

	// Optimization config
	conf := gso.OptimizationConfig{
		MaxTime: 1000,
		N:       50,
		M:       4,
		Diap:    gso.Interval{Min: minX, Max: maxX},
	}

	multistart(conf, f, 100)
}
