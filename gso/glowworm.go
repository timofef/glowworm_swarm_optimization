package gso

// Single glowworm
type glowworm struct {
	coords    []float64
	luciferin float64
	r         float64
}

func initGlowworm(dim int) *glowworm {
	return &glowworm{
		coords:    make([]float64, dim, dim),
		luciferin: 1.0,
		r:         5,
	}
}
