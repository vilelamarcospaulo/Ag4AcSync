package main

import (
	"Ag4AcSync/GeneticAlgorithm"
	"fmt"
	"sync"
)

var ranges [11]float64
var numbers []int
var wg sync.WaitGroup
var mux sync.Mutex

func Exec(n int) {
	var BestIndividual *GeneticAlgorithm.Individual

	for e := 0; e < n; e++ {
		fmt.Println("EXECUÇÃO: ", e+1)

		Ag := GeneticAlgorithm.GeneticAlgorithm{}
		Ag.Init(40, 100, .016, .1, 2)
		Ag.Run()

		grids := Ag.GenerateGrids(5000, true)

		BestIndex := 0
		BestFitnessPercentage := 0.0

		println("INICIANDO AVALIAÇÃO EM 10.000")
		for i := 0; i < 1; i++ {
			Ag.Individuals[i].CalcFitness(grids)

			//fmt.Println(i+1, ":", *Ag.Individuals[i].DNA, "||", Ag.Individuals[i].Fitness, "||", 100*float64(Ag.Individuals[i].Fitness)/5000.0, "%")

			if Ag.Individuals[i].Fitness > Ag.Individuals[BestIndex].Fitness {
				BestIndex = i
			}
		}

		if BestIndividual == nil || BestIndividual.Fitness < Ag.Individuals[BestIndex].Fitness {
			BestIndividual = &Ag.Individuals[BestIndex]
			BestFitnessPercentage = 100 * float64(BestIndividual.Fitness) / 5000.0
		}

		mux.Lock()
		for i := 0; i < 11; i++ {
			if BestFitnessPercentage <= ranges[i] {
				numbers[i]++
				break
			}
		}
		mux.Unlock()
	}

	fmt.Println("WINNER:", BestIndividual.DNA, "||", BestIndividual.Fitness, "||", 100*float64(BestIndividual.Fitness)/5000.0, "%")
	wg.Done()
}

func main() {
	ranges = [11]float64{50.0, 55.0, 60.0, 65.0, 70.0, 75.0, 80.0, 85.0, 90.0, 95.0, 100.0}
	numbers = make([]int, 11)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Exec(20)
	}

	wg.Wait()
	fmt.Println("END")
	fmt.Println(ranges)
	fmt.Println(numbers)
}
