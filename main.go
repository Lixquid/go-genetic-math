package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const epsilon = 0.000001

// Returns the index of the largest float in a slice.
func MaximumIndex(slice []float64) int {
	largestIndex := -1
	largestValue := math.Inf(-1)
	for i := range slice {
		if slice[i] > largestValue {
			largestIndex = i
			largestValue = slice[i]
		}
	}

	return largestIndex
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	target := flag.Float64("target", 200, "The value the genomes are attempting to solve for.")
	populationSize := flag.Int("populationSize", 100, "The number of genomes in a population.")
	mutationRate := flag.Float64("mutationRate", 0.01, "The chance of a gene mutating when breeding.")
	immortalChampion := flag.Bool("immortalChampion", true, "If set, the fittest genome in a population will automatically survive to the next.")
	genomeSize := flag.Int("genomeSize", 7, "The number of genes (operations) in a genome.")
	crossover := flag.Bool("crossover", true, "If set, breeding a new genome will involve two parents with crossover.")
	diffOutput := flag.Bool("diffOutput", true, "If set, champions will only be printed if they're different to the previous generation's champion.")
	colorOutput := flag.Bool("colorOutput", true, "If set, genomes will be printed colourized, with invalid genes printed in red.")
	flag.Parse()

	population := make([]Genome, *populationSize)
	generationCount := 1
	var lastChampion Genome = nil

	// First pass, fill the population will completely random genomes
	for genome := range population {
		population[genome] = make(Genome, *genomeSize)
		for gene := 0; gene < *genomeSize; gene++ {
			population[genome][gene] = uint8(rand.Uint32() & 0xf)
		}
	}

	for {
		// Calculate the fitnesses of the population
		fitnessScores := make([]float64, *populationSize)
		fitnessTotal := 0.0
		for i := range population {
			fitnessScores[i] = population[i].Fitness(*target)
			fitnessTotal += fitnessScores[i]
		}

		// Print out the current champion
		champion := MaximumIndex(fitnessScores)
		if !*diffOutput || lastChampion == nil || !lastChampion.Equal(population[champion]) {
			championValue := population[champion].Value()

			fmt.Printf("Generation %d: %s = %f\n", generationCount, population[champion].String(*colorOutput), championValue)
			lastChampion = population[champion]

			// Check if we've made it to the target
			if *target-epsilon < championValue && championValue < *target+epsilon {
				fmt.Print("Solution found!\n")
				os.Exit(0)
			}
		}

		// Make the next population
		newPopulation := make([]Genome, *populationSize)
		for i := range newPopulation {
			if i == 0 && *immortalChampion {
				newPopulation[0] = population[champion]
			} else {
				newPopulation[i] = Breed(population, fitnessScores, fitnessTotal, *mutationRate, *crossover)
			}
		}

		population = newPopulation
		generationCount++
	}

}
