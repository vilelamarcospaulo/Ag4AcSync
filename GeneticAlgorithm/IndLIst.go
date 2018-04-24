package GeneticAlgorithm

type IndList []Individual

func (list IndList) Len() int {
	return len(list)
}

func (list IndList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list IndList) Less(i, j int) bool {
	return list[i].Fitness > list[j].Fitness
}
