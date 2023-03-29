package gso

func OptimizeSig(conf OptimizationConfig, f Function) ([]float64, float64) {
	// Swarm deployment phase
	s := initSwarm(conf.N, conf.Diap, conf.M)
	// Main cycle
	var (
		maxInd  int
		maxVal  float64
		maxVals []float64
	)
	for t := 0; t < conf.MaxTime; t++ {
		maxInd, maxVal = s.updateLuciferin(f)
		maxVals = append(maxVals, maxVal)
		s.moveGlowwormsSig(t, conf.MaxTime)
		s.updateNeigbourhoodRadius()
		if t > conf.DeltaT {
			if maxVals[t]-maxVals[t-conf.DeltaT] < conf.DeltaF {
				break
			}
		}

	}

	return s.glowworms[maxInd].coords, maxVal
}
