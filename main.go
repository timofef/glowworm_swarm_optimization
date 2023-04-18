package main

import (
	"fmt"
	"glowworm_swarm_optimization/siggso"
	"glowworm_swarm_optimization/test_functions"
	"glowworm_swarm_optimization/utils"
	"math"
	"math/rand"
)

func multistart(conf utils.OptimizationConfig, f test_functions.Function, num int, xMin []float64, yMin float64) {
	max := -math.MaxFloat64
	maxNormX := math.MaxFloat64
	maxNormY := math.MaxFloat64

	sumMaxNormX := 0.0
	sumMaxNormY := 0.0

	//var vec []float64
	succ := 0 // Number of successful launches

	for i := 0; i < num; i++ {
		_, v, val := siggso.Optimize(conf, f)
		if val > max {
			max = val
			//vec = v
		}

		normX := utils.Norm(v, xMin)
		sumMaxNormX += normX
		if normX < maxNormX {
			maxNormX = normX
		}

		normY := math.Abs(val - yMin)
		sumMaxNormY += normY
		if normY < maxNormY {
			maxNormY = normY
		}

		if utils.Norm(v, xMin) < 1e-2 {
			succ++
		}
	}

	fmt.Printf("Min value: %.15f\n", -max)
	//fmt.Printf("Coords:    %v\n", vec)
	fmt.Printf("Best X presision: %.15f\n", maxNormX)
	fmt.Printf("Best Y presision: %.15f\n", maxNormY)
	fmt.Printf("M X presision: %.15f\n", sumMaxNormX/float64(num))
	fmt.Printf("M Y presision: %.15f\n", sumMaxNormY/float64(num))
	fmt.Printf("Success rate: %f\n", float64(succ)/float64(num))
}

func main() {
	rand.Seed(12)
	dims := 2

	// Get test function and diap for x vector
	f, diap, xMin, yMin := test_functions.GetSphere(dims)

	// Optimization config
	conf := utils.OptimizationConfig{
		MaxTime: 1000,
		DeltaF:  1e-6, // Corriodor height
		DeltaT:  20,   // Corridor length
		N:       50,   // Swarm size
		M:       dims, // Dimensions
		Diap:    diap,
	}

	multistart(conf, f, 100, xMin, yMin)
}
