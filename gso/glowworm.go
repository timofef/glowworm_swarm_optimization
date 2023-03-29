package gso

// Single glowworm
type glowworm struct {
	coords    []float64
	val       float64 // Value at coordinates
	luciferin float64 // Luciferin level
	r         float64 // Initial local desision radius
	s         float64 // Initial step size (fixed in basic variation)
}

func initGlowworm(dim int) *glowworm {
	return &glowworm{
		coords:    make([]float64, dim, dim),
		luciferin: 1.0,
		r:         5.0,
		s:         0.2,
	}
}
