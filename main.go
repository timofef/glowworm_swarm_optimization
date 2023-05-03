package main

import (
	"fmt"
	"glowworm_swarm_optimization/siggso"
	"glowworm_swarm_optimization/test_functions"
	"glowworm_swarm_optimization/utils"
	"math"
	"math/rand"
)

func multistart(conf utils.OptimizationConfig, f test_functions.Function, diap utils.Interval, num int, xMin []float64, yMin float64, r *rand.Rand) {
	max := -math.MaxFloat64
	maxNormX := math.MaxFloat64
	maxNormY := math.MaxFloat64

	epsX := 0.01 * (diap.Max - diap.Min) * math.Sqrt(float64(len(xMin)))

	sumMaxNormX := 0.0
	sumMaxNormY := 0.0

	//var vec []float64
	succ := 0 // Number of successful launches

	/*res := make([]float64, 0)*/

	for i := 0; i < num; i++ {
		_, v, val, _ := siggso.Optimize(conf, f, r)

		/*res = append(res, vals...)*/

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

		if utils.Norm(v, xMin) <= epsX {
			succ++
		}
	}

	// Write to file
	/*file, _ := os.Create("res_siggso.csv")
	for j := 0; j < len(res); j++ {
		fmt.Fprintf(file, "%e\n", res[j])
	}
	file.Close()*/

	fmt.Printf("Min value: %e\n", -max)
	//fmt.Printf("Coords:    %v\n", vec)
	fmt.Printf("Best X presision: %e\n", maxNormX)
	fmt.Printf("Best Y presision: %e\n", maxNormY)
	fmt.Printf("M X presision: %e\n", sumMaxNormX/float64(num))
	fmt.Printf("M Y presision: %e\n", sumMaxNormY/float64(num))
	fmt.Printf("Success rate: %.2f\n", float64(succ)/float64(num))
	fmt.Println()
}

func main() {
	r := rand.New(rand.NewSource(12))

	dims := []int{2, 4, 8, 16, 32, 64}

	for _, dim := range dims {
		// Get test function and diap for x vector
		f, diap, xMin, yMin := test_functions.GetSphere(dim)

		// Optimization config
		conf := utils.OptimizationConfig{
			MaxTime: 1000,
			DeltaF:  1e-6, // Corriodor height
			DeltaT:  1000, // Corridor length
			N:       50,   // Swarm size
			M:       dim,  // Dimensions
			Diap:    diap,
		}

		multistart(conf, f, diap, 100, xMin, yMin, r)
	}
}
