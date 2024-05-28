package pkg

import (
	"gonum.org/v1/gonum/mat"
)

type DropMatrix struct {
	Description string
	Matrix      *mat.Dense
}

type DropOdds interface {
	GetWantedDrops() [][]DropGoal
	GetName() string
	// Calculates the transition matrices for every wanted drop list
	GetTransitionMatrices() []DropMatrix

	// Check whether every drop on the wanted drop list in at least one of the droptables.
	IsPossible() bool
}
