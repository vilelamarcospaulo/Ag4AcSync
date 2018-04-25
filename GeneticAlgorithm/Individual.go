package GeneticAlgorithm

import (
	"CellularAutomaton/ac"
	"math"
	"math/rand"
)

type Individual struct {
	DNA     *([]byte)
	Fitness int
	Radius  int
}

func (ind *Individual) Init(radius int, position int) {
	size := int(math.Pow(2, float64(2*radius+1)))
	dna := make([]byte, size)

	ind.DNA = &dna
	ind.Radius = radius
	ind.Fitness = -1

	if position < 0 {
		return
	}

	for i := 0; i < size; i++ {
		dice := rand.Intn(100)
		if dice < position {
			dna[i] = 1
		} else {
			dna[i] = 0
		}
	}
}

func toInteger(a []byte) int {
	length := len(a)
	sum := 0
	for i := 0; i < length; i++ {
		value := 0
		if a[length-i-1] == 1 {
			value = 1
		}
		j := float64(i)
		sum += value * int(math.Pow(2, j))
	}

	return sum
}

func CheckAlternate(CellularAutomaton ac.CelullarAutomaton) int {
	pprevius := CellularAutomaton.Grid.PPCells
	previus := CellularAutomaton.Grid.PCells
	last := CellularAutomaton.Grid.Cells

	pattern := last[0]
	for _, b := range last {
		if b != pattern {
			return 0
		}
	}

	for _, b := range pprevius {
		if b != pattern {
			return 0
		}
	}

	if pattern == 0 {
		pattern = 1
	} else {
		pattern = 0
	}

	for _, b := range previus {
		if b != pattern {
			return 0
		}
	}

	return 1
}

func (ind *Individual) CalcFitness(grid [][]byte) int {
	// ind.Fitness = toInteger(*ind.DNA)
	// return ind.Fitness

	ind.Fitness = 0
	for i := 0; i < len(grid); i++ {
		CellularAutomaton := ac.Create(300, grid[i], *ind.DNA, ind.Radius)
		CellularAutomaton.Run()

		ind.Fitness += CheckAlternate(CellularAutomaton)
	}

	return ind.Fitness
}

func (ind *Individual) Mutation(probability float32) {
	for i := 0; i < len(*ind.DNA); i++ {
		dice := rand.Float32()
		if dice <= probability {
			dice = rand.Float32()
			if dice <= .5 {
				(*ind.DNA)[i] = 1
			} else {
				(*ind.DNA)[i] = 0
			}
		}
	}
}
