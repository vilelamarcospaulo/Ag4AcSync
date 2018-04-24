package GeneticAlgorithm

import (
	"math/rand"
)

func (ag GeneticAlgorithm) SelectRandomElit() int {
	index := rand.Intn(ag.ElitismSize)

	return index
}
