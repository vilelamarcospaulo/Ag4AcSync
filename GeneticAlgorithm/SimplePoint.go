package GeneticAlgorithm

import (
	"math/rand"
)

type SimplePoint struct {
	point int
}

func (sp *SimplePoint) init(size int) {
	sp.point = rand.Intn(size)
}

func (sp SimplePoint) Execute(father1 Individual, father2 Individual, son1 *Individual, son2 *Individual) {
	size := len(*father1.DNA)
	sp.init(size)

	for i := 0; i <= sp.point; i++ {
		(*son1.DNA)[i] = (*father1.DNA)[i]
		(*son2.DNA)[i] = (*father2.DNA)[i]
	}

	for i := sp.point; i < size; i++ {
		(*son1.DNA)[i] = (*father2.DNA)[i]
		(*son2.DNA)[i] = (*father1.DNA)[i]
	}
}
