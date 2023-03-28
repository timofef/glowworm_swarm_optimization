package gso

import (
	"math"
	"math/rand"
)

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

func (s *swarm) updateLuciferin(f Function) {
	for i := range s.glowworms {
		decay := (1. - s.ro) * s.glowworms[i].luciferin
		increase := s.gamma * f(s.glowworms[i].coords)
		s.glowworms[i].luciferin = decay + increase
	}
}

func (s *swarm) moveGlowworms() {
	directions := make([][]float64, len(s.glowworms), len(s.glowworms))

	for j, worm := range s.glowworms {
		// For each glowworm find neighbours
		neighbours, sum := s.getNeighbours(j)
		// Best worm doesn't move
		if len(neighbours) != 0 {
			// For each neighbour find probability of moving towards it
			probs := make([]float64, len(neighbours), len(neighbours))
			for i := range neighbours {
				probs[i] = (neighbours[i].luciferin - worm.luciferin) / sum
			}
			// Choose direction based on probability
			toss := rand.Float64()
			lower, upper := 0.0, 0.0
			var best *glowworm = nil
			for i := 0; i < len(probs); i++ {
				lower = upper
				upper = lower + probs[i]
				if toss >= lower && toss <= upper {
					best = neighbours[i]
					break
				}
			}
			n := norm(best.coords, worm.coords)
			// Calculate shift in chosen direction
			directions[j] = make([]float64, len(worm.coords), len(worm.coords))
			for i := 0; i < len(worm.coords); i++ {
				directions[j][i] = worm.coords[i] + worm.s*(best.coords[i]-worm.coords[i])/n
			}
		} else {
			directions[j] = make([]float64, len(worm.coords), len(worm.coords))
			for i := 0; i < len(worm.coords); i++ {
				directions[j][i] = worm.coords[i]
			}
		}
	}
	// Move each glowworm
	for j := range s.glowworms {
		s.glowworms[j].coords = directions[j]
	}
}

func (s *swarm) getNeighbours(index int) ([]*glowworm, float64) {
	var neighbours []*glowworm
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
