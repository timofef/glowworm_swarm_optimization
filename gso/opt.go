package gso

type Interval struct {
	Min float64
	Max float64
}

type OptimizationConfig struct {
	MaxTime int // Time steps limit
	S       int // Swarm size
	Diap    Interval
}

type Function struct {
	N int                          // Space dimensions
	F func(int, []float64) float64 // Target function
}

func Optimize(conf OptimizationConfig, f Function) float64 {
	// Create swarm
	s := initSwarm(conf.S, conf.Diap, f.N)

	for t := 0; t < conf.MaxTime; t++ {
		s.updateLuciferin(f)

		// Move glowworms
		
		// Update decision radius
	}

	return 0
}
