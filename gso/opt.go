package gso

import (
	"glowworm_swarm_optimization/test_functions"
	"glowworm_swarm_optimization/utils"
	"math"
	"math/rand"
)

const (
	RO    = 0.4 // Luciferin decay factor
	GAMMA = 0.6 // Fitness enhancement factor

	L0 = 5.0 // Initial luciferin level
	R0 = 5   // Initial local decision radius
	S  = 0.5 // Step size

	RS = 500.0 // Sensor range of the glowworms

	NT   = 50   // Threshold value of the number of glowworms
	BETA = 0.08 // Domain change rate
)

func Optimize(conf utils.OptimizationConfig, f test_functions.Function, r *rand.Rand) ([]*glowworm, []float64, float64, []float64) {
	// Swarm deployment phase
	s := initSwarm(conf.N, conf.Diap, conf.M, r)

	// Main cycle
	var (
		maxInd int
		maxVal float64 = math.MaxFloat64
		//lastMaxVal float64
		maxVals []float64
		//stagnation int
	)

	for t := 0; t < conf.MaxTime; t++ {
		//lastMaxVal = maxVal
		maxInd, maxVal = s.updateLuciferin(f)
		maxVals = append(maxVals, maxVal)

		s.moveGlowworms(r)
		s.updateNeigbourhoodRadius()

		/*if lastMaxVal-maxVal > conf.DeltaF {
			stagnation++
			if stagnation == conf.DeltaT {
				break
			}
		} else {
			stagnation = 0
		}*/

		if t > conf.DeltaT {
			if math.Abs(maxVals[t]-maxVals[t-conf.DeltaT]) < conf.DeltaF {
				break
			}
		}
	}

	return s.glowworms, s.glowworms[maxInd].Coords, maxVal, maxVals
}
