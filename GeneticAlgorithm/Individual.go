package GeneticAlgorithm

type Individual struct {
	DNA     *([]byte)
	Fitness int
}

func CreateIndividual(radius int) Individual {
	size := 2*radius + 1
	dna := make([]byte, size)

	for i := 0; i < size; i++ {
		dna[i] = 1 // TODO :: RAND VALUE
	}

	i := Individual{
		DNA:     &dna,
		Fitness: -1,
	}
}

func (i Individual) CalcFitness() int {
	if i.Fitness > -1 {
		return i.Fitness
	}

	// TODO :: RODAR O DNA NOS ACS RECEBIDOS

	i.Fitness = 0
	return i.Fitness
}

func (i Individual) Mutation() {
	// TODO :: MUTAÇÃO NO DNA
}
