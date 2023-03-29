package gso

import (
	"math"
	"math/rand"
)

type Interval struct {
	Min float64
	Max float64
}

type OptimizationConfig struct {
	MaxTime int      // Time steps limit
	DeltaF  float64  // Corridor width
	DeltaT  int      // Corridor length
	N       int      // Swarm size
	M       int      // Dimension of solution
	Diap    Interval // Diap for x vector
}

type Function func([]float64) float64

// Glowworm swarm
type swarm struct {
	glowworms []*glowworm
	ro        float64 // Luciferin decay factor
	gamma     float64 // Fitness enhancement factor
}

func initSwarm(swarmSize int, diap Interval, dim int) *swarm {
	glowworms := make([]*glowworm, swarmSize, swarmSize)
	s := &swarm{
		glowworms: glowworms,
		ro:        0.4,
		gamma:     0.6,
	}
	// Init coords with random values in interval
	for i := 0; i < swarmSize; i++ {
		g := initGlowworm(dim)
		for j := 0; j < dim; j++ {
			g.coords[j] = diap.Min + rand.Float64()*(diap.Max-diap.Min)
		}
		s.glowworms[i] = g
	}

	return s
}

func (s *swarm) updateLuciferin(f Function) (int, float64) {
	// This stage is combined with calculating fitness function and determining best worm
	maxVal := -1e30
	maxInd := 0

	for i := range s.glowworms {
		s.glowworms[i].val = f(s.glowworms[i].coords)
		if s.glowworms[i].val > maxVal {
			maxVal = s.glowworms[i].val
			maxInd = i
		}

		decay := (1. - s.ro) * s.glowworms[i].luciferin
		increase := s.gamma * s.glowworms[i].val
		s.glowworms[i].luciferin = decay + increase
	}

	return maxInd, maxVal
}

func (s *swarm) getNeighbours(index int) ([]*glowworm, float64) {
	var neighbours []*glowworm
	// Sum is needed to calculate probability of movement in direction of neighbour
	sum := 0.0
	for _, potentialNeighbour := range s.glowworms {
		if potentialNeighbour != s.glowworms[index] {
			if potentialNeighbour.luciferin > s.glowworms[index].luciferin {
				n := norm(s.glowworms[index].coords, potentialNeighbour.coords)
				if n < s.glowworms[index].r {
					neighbours = append(neighbours, potentialNeighbour)
					sum += potentialNeighbour.luciferin - s.glowworms[index].luciferin
				}
			}
		}
	}
	return neighbours, sum
}

func (s *swarm) updateNeigbourhoodRadius() {
	for i := 0; i < len(s.glowworms); i++ {
		neighbours, _ := s.getNeighbours(i)
		s.glowworms[i].r = math.Min(5, math.Max(0, s.glowworms[i].r+0.08*(5.0-float64(len(neighbours)))))
	}
}

func norm(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += math.Pow(a[i]-b[i], 2.0)
	}
	return math.Sqrt(sum)
}
