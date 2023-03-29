package main

import (
	"fmt"
	"glowworm_swarm_optimization/gso"
	"glowworm_swarm_optimization/test_functions"
	"math/rand"
)

func multistart(conf gso.OptimizationConfig, f gso.Function, num int, xMin []float64, yMin float64) {
	max := -1e30
	var vec []float64

	//succProb := 0.0

	for i := 0; i < num; i++ {
		v, val := gso.OptimizeSig(conf, f)
		if val > max {
			max = val
			vec = v
		}
	}

	fmt.Printf("Min value: %f\n", -max)
	fmt.Printf("Coords:    %v\n", vec)
	//fmt.Printf("Success rate: %f\n", succProb)
}

func main() {
	rand.Seed(12345)
	dims := 2

	// Get test function and diap for x vector
	f, diap, xMin, yMin := test_functions.GetRosenbrok(dims)

	// Optimization config
	conf := gso.OptimizationConfig{
		MaxTime: 1000,
		DeltaF:  1e-6, // Corriodor height
		DeltaT:  20,   // Corridor length
		N:       50,   // Swarm size
		M:       dims, // Dimensions
		Diap:    diap,
	}

	multistart(conf, f, 100, xMin, yMin)
}
