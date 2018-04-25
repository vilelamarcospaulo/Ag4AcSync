package GeneticAlgorithm

import (
	"fmt"
	"math/rand"
	"sort"
)

type GeneticAlgorithm struct {
	CurrentGeneration   int
	NumberGenerations   int
	PopulationSize      int
	NumberChildrens     int
	MutationProbability float32
	ElitismSize         int
	Radius              int

	Individuals []Individual
	Childrens   []Individual
}

func (ag *GeneticAlgorithm) Init(generations int, population int, mutationProbability float32, elitismPercentage float32, radius int) {
	ag.CurrentGeneration = 0

	ag.ElitismSize = int(elitismPercentage * float32(population))
	ag.NumberGenerations = generations
	ag.PopulationSize = population
	ag.NumberChildrens = ag.PopulationSize - ag.ElitismSize
	ag.MutationProbability = mutationProbability

	ag.Individuals = make([]Individual, ag.PopulationSize) //TODO :: RANDOM POPULATION
	ag.Childrens = make([]Individual, ag.NumberChildrens)

	fmt.Println("GERANDO POPULACAO INICIAL")
	grids := ag.GenerateGrids(100, false)

	for i := 0; i < ag.PopulationSize; i++ {
		ag.Individuals[i].Init(radius, i)
		ag.Individuals[i].CalcFitness(grids)

		// fmt.Println("CALCULANDO FITNESS ELEMENTO: ", i, "||")
		// fmt.Println("||", ag.Individuals[i].CalcFitness(grids))
	}
	sort.Sort(IndList(ag.Individuals))

	for i := 0; i < ag.NumberChildrens; i++ {
		ag.Childrens[i].Init(radius, -1)
	}

}

func Distrib(percentOfOne int) byte {
	if rand.Intn(100) <= percentOfOne {
		return 1
	}
	return 0
}

func (ag *GeneticAlgorithm) GenerateGrids(size int, random bool) [][]byte {
	grids := make([][]byte, size)
	for i := 0; i < size; i++ {
		grids[i] = make([]byte, 149)

		for j := 0; j < 149; j++ {
			if random {
				grids[i][j] = Distrib(50)
				continue
			}
			grids[i][j] = Distrib(i)
		}
	}

	return grids
}

func (ag *GeneticAlgorithm) GenerateSons(i int, j int, grids [][]byte) {
	var father1, father2 Individual
	var index1, index2 int

	index1 = ag.SelectRandomElit()
	father1 = ag.Individuals[index1]
	//fmt.Println("GERACAO:", ag.CurrentGeneration, "PAI 1:", index1, "||", *father1.DNA, "||", father1.Fitness)

	for index2 = ag.SelectRandomElit(); index1 == index2; index2 = ag.SelectRandomElit() {
	}

	father2 = ag.Individuals[index2]
	//fmt.Println("GERACAO:", ag.CurrentGeneration, "PAI 2:", index2, "||", *father2.DNA, "||", father2.Fitness)

	son1 := &ag.Childrens[i]
	son2 := &ag.Childrens[j]

	croosover := SimplePoint{}
	croosover.Execute(father1, father2, son1, son2)

	son1.Mutation(ag.MutationProbability)
	son2.Mutation(ag.MutationProbability)

	//fmt.Print("GERACAO: ", ag.CurrentGeneration, " CALCULANDO FITNESS FILHO: ", i, "||", *son1.DNA)
	son1.CalcFitness(grids)
	//fmt.Println("||", son1.Fitness)

	//fmt.Print("GERACAO: ", ag.CurrentGeneration, " CALCULANDO FITNESS FILHO: ", j, "||", *son2.DNA)
	son2.CalcFitness(grids)
	//fmt.Println("||", son2.Fitness)

	// xxfmt.Println(father1.DNA, father1.Fitness)
	// fmt.Println(father2.DNA, father2.Fitness)
	// fmt.Println(son1.DNA, son1.Fitness)
	// fmt.Println(son2.DNA, son2.Fitness)
	// fmt.Println("---")
}

func (ag *GeneticAlgorithm) NextGeneration() {
	ag.CurrentGeneration++

	grids := ag.GenerateGrids(100, false)

	fmt.Println("GERACAO:", ag.CurrentGeneration)

	for i := 0; i < ag.NumberChildrens; i += 2 {
		ag.GenerateSons(i, i+1, grids)
	}

	ag.Reinsertion()
}

func (ag *GeneticAlgorithm) Reinsertion() {
	sort.Sort(IndList(ag.Individuals))
	sort.Sort(IndList(ag.Childrens))

	//fmt.Println("MELHOR FILHO GERADO NA GERACAO:", ag.Childrens[0].DNA, "||", ag.Childrens[0].Fitness)

	for i := ag.ElitismSize; i < ag.PopulationSize; i++ {
		ag.Individuals[i], ag.Childrens[i-ag.ElitismSize] = ag.Childrens[i-ag.ElitismSize], ag.Individuals[i]
	}

	sort.Sort(IndList(ag.Individuals))
}

func (ag *GeneticAlgorithm) Run() {
	for ag.CurrentGeneration < ag.NumberGenerations {
		ag.NextGeneration()
	}
}
