package gso

// Single glowworm
type glowworm struct {
	coords    []float64
	val       float64 // Value at coordinates
	luciferin float64 // Luciferin level
	r         float64 // Local desision radius
	s         float64 // Fixed step size
}

func initGlowworm(dim int) *glowworm {
	return &glowworm{
		coords:    make([]float64, dim, dim),
		luciferin: 1.0,
		r:         5.0,
		s:         0.2,
	}
}
