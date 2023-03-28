package main

import (
	"fmt"
	"glowworm_swarm_optimization/gso"
	"glowworm_swarm_optimization/test_functions"
	"math/rand"
)

func main() {
	rand.Seed(69420)

	// Get test function and diap for x vector
	sphere, minX, maxX := test_functions.GetSphere(2)

	// Optimization config
	conf := gso.OptimizationConfig{
		MaxTime: 1000,
		N:       50,
		M:       64,
		Diap:    gso.Interval{Min: minX, Max: maxX},
	}

	min := -1000000.0

	for i := 0; i < 1; i++ {
		val := gso.Optimize(conf, sphere)
		if val > min {
			min = val
		}
	}

	/*min := gso.Optimize(conf, sphere)*/

	fmt.Printf("Min: %f\n", -min)
}
