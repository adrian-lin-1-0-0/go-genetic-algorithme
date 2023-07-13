package ga

type Organism interface {
	GetChromosome() Chromosome
	Fitness() float64
}

type OrganismFactory interface {
	Create() Organism
	CreateWithChromosome(Chromosome) Organism
	CreateGene() Gene
}

type Population []Organism

type Chromosome []Gene

type Gene rune
