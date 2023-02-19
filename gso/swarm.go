package gso

import (
	"math"
	"math/rand"
)

// Glowworm swarm
type swarm struct {
	glowworms []*glowworm
	ro        float64
	gamma     float64
}

func initSwarm(n int, diap Interval, dim int) *swarm {
	glowworms := make([]*glowworm, n, n)
	s := &swarm{
		glowworms: glowworms,
		ro:        0.4,
		gamma:     0.6,
	}

	// Init coords with random values in interval
	for i := 0; i < n; i++ {
		g := initGlowworm(dim)
		for j := 0; j < dim; j++ {
			g.coords[j] = diap.Min + rand.Float64()*(diap.Max-diap.Min)
		}
		s.glowworms[i] = g
	}

	return s
}

func (s *swarm) updateLuciferin(f Function) {
	for i := 0; i < len(s.glowworms); i++ {
		decay := (1. - s.ro) * s.glowworms[i].luciferin
		increase := s.gamma * f.F(f.N, s.glowworms[i].coords)
		s.glowworms[i].luciferin = decay + increase
	}
}

func (s *swarm) move() {

}

func norm(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += math.Pow(a[i]-b[i], 2.0)
	}
	return math.Sqrt(sum)
}
