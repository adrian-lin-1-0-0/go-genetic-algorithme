package ga

type GeneticAlgorithmOptions struct {
	MaxGeneration    int
	MutationRate     float64
	OrganismFactory  OrganismFactory
	PopulationSize   int
	FitnessThreshold float64
}

func (opts *GeneticAlgorithmOptions) SetMaxGeneration(maxGeneration int) *GeneticAlgorithmOptions {
	opts.MaxGeneration = maxGeneration
	return opts
}

func (opts *GeneticAlgorithmOptions) SetMutationRate(mutationRate float64) *GeneticAlgorithmOptions {
	opts.MutationRate = mutationRate
	return opts
}

func (opts *GeneticAlgorithmOptions) SetOrganismFactory(organismFactory OrganismFactory) *GeneticAlgorithmOptions {
	opts.OrganismFactory = organismFactory
	return opts
}

func (opts *GeneticAlgorithmOptions) SetPopulationSize(populationSize int) *GeneticAlgorithmOptions {
	opts.PopulationSize = populationSize
	return opts
}

func (opts *GeneticAlgorithmOptions) SetFitnessThreshold(fitnessThreshold float64) *GeneticAlgorithmOptions {
	opts.FitnessThreshold = fitnessThreshold
	return opts
}
