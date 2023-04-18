package siggso

type glowworm struct {
	Coords     []float64
	Val        float64 // Value at coordinates
	luciferin  float64 // Luciferin level
	r          float64 // Local decision radius
	s          float64 // Step size
	neighbours int     // Neighbours in neighbourhood
}

func initGlowworm(dim int) *glowworm {
	return &glowworm{
		Coords:    make([]float64, dim, dim),
		luciferin: L0,
		r:         R0,
		s:         S,
	}
}
