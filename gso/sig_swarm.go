package gso

import (
	"math"
	"math/rand"
)

func (s *swarm) moveGlowwormsSig(t, tMax int) {
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
				directions[j][i] = worm.coords[i] + worm.s*fi(t, tMax)*(best.coords[i]-worm.coords[i])/n
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

func fi(t, tMax int) float64 {
	return 1.0 / (1.0 + math.Exp(-50.0*(float64(t)/float64(tMax)-0.1)))
}
