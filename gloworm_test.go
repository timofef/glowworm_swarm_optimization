package main

import (
	"glowworm_swarm_optimization/gso"
	"glowworm_swarm_optimization/test_functions"
	"math/rand"
	"testing"
)

func TestGloworm(t *testing.T) {
	rand.Seed(12)
	dims := 2

	// Get test function and diap for x vector
	f, diap, xMin, yMin := test_functions.GetSphere(dims)

	// Optimization config
	conf := gso.OptimizationConfig{
		MaxTime: 1000,
		DeltaF:  1e-6, // Corriodor height
		DeltaT:  20,   // Corridor length
		N:       50,   // Swarm size
		M:       dims, // Dimensions
		Diap:    diap,
	}

	multistart(conf, f, 100, xMin, yMin)
}
