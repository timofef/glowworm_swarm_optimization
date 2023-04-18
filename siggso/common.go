package siggso

import (
	"glowworm_swarm_optimization/test_functions"
	"glowworm_swarm_optimization/utils"
	"math"
	"math/rand"
)

type swarm struct {
	glowworms []*glowworm
	ro        float64 // Luciferin decay factor
	gamma     float64 // Fitness enhancement factor
}

func initSwarm(swarmSize int, diap utils.Interval, dim int) *swarm {
	glowworms := make([]*glowworm, swarmSize, swarmSize)
	s := &swarm{
		glowworms: glowworms,
		ro:        RO,
		gamma:     GAMMA,
	}
	// Init Coords with random values in interval
	for i := 0; i < swarmSize; i++ {
		g := initGlowworm(dim)
		for j := 0; j < dim; j++ {
			g.Coords[j] = diap.Min + rand.Float64()*(diap.Max-diap.Min)
		}
		s.glowworms[i] = g
	}

	return s
}

func (s *swarm) updateLuciferin(f test_functions.Function) (int, float64) {
	// This stage is combined with calculating fitness function
	//and determining best worm
	maxVal := -math.MaxFloat64
	maxInd := 0

	for i := range s.glowworms {
		s.glowworms[i].Val = f(s.glowworms[i].Coords)
		if s.glowworms[i].Val > maxVal {
			maxVal = s.glowworms[i].Val
			maxInd = i
		}

		decay := (1. - s.ro) * s.glowworms[i].luciferin
		increase := s.gamma * s.glowworms[i].Val
		s.glowworms[i].luciferin = decay + increase
	}

	return maxInd, maxVal
}

func (s *swarm) getNeighbours(index int) ([]*glowworm, float64) {
	var neighbours []*glowworm
	// Sum is needed to calculate probability of movement in direction of neighbour
	sum := 0.0
	for i, potentialNeighbour := range s.glowworms {
		if potentialNeighbour != s.glowworms[index] {
			if potentialNeighbour.luciferin > s.glowworms[index].luciferin {
				n := utils.Norm(s.glowworms[index].Coords, potentialNeighbour.Coords)
				if n < s.glowworms[index].r {
					neighbours = append(neighbours, potentialNeighbour)
					sum += potentialNeighbour.luciferin - s.glowworms[index].luciferin
					s.glowworms[i].neighbours = len(neighbours)
				}
			}
		}
	}
	return neighbours, sum
}

func (s *swarm) updateNeigbourhoodRadius() {
	for i := 0; i < len(s.glowworms); i++ {
		s.glowworms[i].r = math.Min(RS, math.Max(0, s.glowworms[i].r+BETA*(float64(NT)-float64(s.glowworms[i].neighbours))))
	}
}

// Sigmoid
func fi(t, maxT int) float64 {
	return 1. - 1./(1.+math.Exp(-LAMBDA*(float64(t)/float64(maxT)-EPS)))
}
