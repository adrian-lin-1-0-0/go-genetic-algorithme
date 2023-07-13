package ga

import (
	"math/rand"
	"time"
)

type GeneticAlgorithm struct {
	maxGeneration    int
	mutationRate     float64
	organismFactory  OrganismFactory
	populationSize   int
	bestOrganism     Organism
	fitnessThreshold float64
	rand             *rand.Rand
}

func NewGeneticAlgorithm(options *GeneticAlgorithmOptions) *GeneticAlgorithm {
	if options.MaxGeneration == 0 {
		options.MaxGeneration = 100
	}

	return &GeneticAlgorithm{
		maxGeneration:    options.MaxGeneration,
		mutationRate:     options.MutationRate,
		organismFactory:  options.OrganismFactory,
		populationSize:   options.PopulationSize,
		fitnessThreshold: options.FitnessThreshold,
		rand:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ga *GeneticAlgorithm) Run() Organism {
	currentPopulation := ga.createPopulation()
	for generation := 0; generation < ga.maxGeneration; generation++ {
		currentPopulation = pipe(currentPopulation,
			ga.selection,
			ga.crossover,
			ga.mutate)
		ga.bestOrganism = ga.findBestOrganism(currentPopulation)
		if ga.terminationCondition() {
			break
		}
	}
	return ga.bestOrganism
}

func (ga *GeneticAlgorithm) createPopulation() Population {
	population := make(Population, 0)
	for i := 0; i < ga.populationSize; i++ {
		population = append(population, ga.organismFactory.Create())
	}
	return population
}

func (ga *GeneticAlgorithm) selection(population Population) Population {
	totalFitness := 0.0

	for _, organism := range population {
		totalFitness += organism.Fitness()
	}

	probability := make([]float64, len(population))
	for i, organism := range population {
		probability[i] = organism.Fitness() / totalFitness
	}

	cumulativeProbability := make([]float64, len(population))
	cumulativeProbability[0] = probability[0]
	for i := 1; i < len(population); i++ {
		cumulativeProbability[i] = cumulativeProbability[i-1] + probability[i]
	}

	newPopulation := make(Population, 0)
	for i := 0; i < len(population); i++ {
		r := ga.rand.Float64()
		for j, probability := range cumulativeProbability {
			if r <= probability {
				newPopulation = append(newPopulation, population[j])
				break
			}
		}
	}
	return newPopulation
}

func (ga *GeneticAlgorithm) crossover(population Population) Population {

	newPopulation := make(Population, 0)
	for i := 0; i < len(population); i++ {
		a := population[ga.rand.Intn(len(population))]
		b := population[ga.rand.Intn(len(population))]
		child := ga.crossoverOrganism(a, b)
		newPopulation = append(newPopulation, child)
	}

	return newPopulation
}

func (ga *GeneticAlgorithm) crossoverOrganism(a, b Organism) Organism {
	aChromosome := a.GetChromosome()
	bChromosome := b.GetChromosome()

	crossoverPoint := ga.rand.Intn(len(aChromosome))

	childChromosome := append(aChromosome[:crossoverPoint:crossoverPoint], bChromosome[crossoverPoint:]...)
	child := ga.organismFactory.CreateWithChromosome(childChromosome)

	return child
}

func (ga *GeneticAlgorithm) mutate(population Population) Population {

	mutatePopulation := make(Population, 0)
	for _, organism := range population {
		chromosome := organism.GetChromosome()
		for i := 0; i < len(chromosome); i++ {
			if ga.rand.Float64() < ga.mutationRate {
				chromosome[i] = ga.organismFactory.CreateGene()
			}
		}
		mutatePopulation = append(mutatePopulation, ga.organismFactory.CreateWithChromosome(chromosome))
	}

	return mutatePopulation
}

func (ga *GeneticAlgorithm) terminationCondition() bool {
	return ga.bestOrganism.Fitness() >= ga.fitnessThreshold
}

func (ga *GeneticAlgorithm) findBestOrganism(population Population) Organism {
	best := 0.0
	bestIdx := 0
	for idx, organism := range population {
		if organism.Fitness() > best {
			best = organism.Fitness()
			bestIdx = idx
		}
	}

	return population[bestIdx]
}
