package GeneticAlgorithm

type GeneticAlgorithm struct {
	CurrentGeneration   int
	NumberGenerations   int
	PopulationSize      int
	NumberChildrens     int
	MutationProbability float32

	Individuals []Individual
	Childrens   []Individual
}

func (ag GeneticAlgorithm) init(generations int, population int, childrens int, mutation float32) {
	ag.CurrentGeneration = 0

	ag.NumberGenerations = generations
	ag.PopulationSize = population
	ag.NumberChildrens = childrens
	ag.MutationProbability = mutation

	ag.Individuals = make([]Individual, ag.PopulationSize) //TODO :: RANDOM POPULATION
	ag.Childrens = make([]Individual, ag.NumberChildrens)
}

func (ag GeneticAlgorithm) NextGeneration() {
	ag.CurrentGeneration++

}

func (ag GeneticAlgorithm) Reinsertion() {

}

func (ag GeneticAlgorithm) Run() {
	for ag.CurrentGeneration < ag.NumberGenerations {

	}
}
