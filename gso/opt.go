package gso

type Interval struct {
	Min float64
	Max float64
}

type OptimizationConfig struct {
	MaxTime int      // Time steps limit
	N       int      // Swarm size
	M       int      // Dimension of solution
	Diap    Interval // Diap for x vector
}

type Function func([]float64) float64

func Optimize(conf OptimizationConfig, f Function) ([]float64, float64) {
	// Swarm deployment phase
	s := initSwarm(conf.N, conf.Diap, conf.M)
	// Main cycle
	for t := 0; t < conf.MaxTime; t++ {
		// Luciferin update phase
		s.updateLuciferin(f)
		// Location update phase
		s.moveGlowworms()
		// Neighbourhood range update phase
		s.updateNeigbourhoodRadius()
	}

	// Find max value and it's coordinates
	max := f(s.glowworms[0].coords)
	ind := 0
	for i := 1; i < len(s.glowworms); i++ {
		val := f(s.glowworms[i].coords)
		if val > max {
			max = val
			ind = i
		}
	}

	return s.glowworms[ind].coords, max
}
