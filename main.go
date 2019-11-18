package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Valid instructions: 0-9 number, 10: addition, 11: subtraction, 12: multiplication, 13: division
type genome []uint8

func (g genome) String() string {
	var output strings.Builder
	for i := range g {
		switch g[i] {
		case 10:
			output.WriteRune('+')
		case 11:
			output.WriteRune('-')
		case 12:
			output.WriteRune('*')
		case 13:
			output.WriteRune('/')
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			output.WriteRune(rune(g[i] + 48))
		default:
			// Invalid gene
			output.WriteRune('X')
		}
	}
	return output.String()
}

const (
	none     = iota
	add      = iota
	subtract = iota
	multiply = iota
	divide   = iota
)

func (g genome) Value() float64 {
	initialized := false
	value := 0.0
	operand := none

	for i := range g {
		switch g[i] {
		case 10:
			if initialized && operand == none {
				operand = add
			}
		case 11:
			if initialized && operand == none {
				operand = subtract
			}
		case 12:
			if initialized && operand == none {
				operand = multiply
			}
		case 13:
			if initialized && operand == none {
				operand = divide
			}
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			if !initialized {
				initialized = true
				value = float64(g[i])
				break
			}
			switch operand {
			case add:
				value += float64(g[i])
			case subtract:
				value -= float64(g[i])
			case multiply:
				value *= float64(g[i])
			case divide:
				value /= float64(g[i])
			}
			operand = none
		}
	}

	return value
}

// Single parent breeding
func Breed(population []genome, fitness []float64, fitnessTotal float64, mutationRate float64) genome {
	fitnessGoal := rand.Float64() * fitnessTotal
	total := 0.0
	for i := range population {
		if total + fitness[i] < fitnessGoal {
			total += fitness[i]
			continue
		}

		new := make(genome, len(population[0]))
		copy(new, population[i])

		// mutate
		// Ideally we'd flip one bit, but it doesn't matter too much
		for j := range new {
			if rand.Float64() < mutationRate {
				new[j] = uint8(rand.Uint32() & 0xf)
			}
		}

		return new
	}

	// Should never ever happen
	panic("Reached end of population when breeding.")
}

func LargestIndex(data []float64) int {
	largestIndex := -1
	largestValue := -1.0
	for i := range data {
		if data[i] > largestValue {
			largestIndex = i
			largestValue = data[i]
		}
	}
	return largestIndex
}

var epsilon = math.Nextafter(1.0, 2.0) - 1.0

func (g genome) Fitness(target int) float64 {
	return 1 / math.Max(math.Abs(float64(target) - g.Value()), epsilon)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	popSize := flag.Int("popsize", 100, "The number of members in a population.")
	mutRate := flag.Float64("mutrate", 0.05, "The mutation rate.")
	immortalChampion := flag.Bool("immortalChampion", false, "If set, the fittest member will automatically survive.")
	genomeSize := flag.Int("genomeSize", 7, "The number of instructions in the genome.")
	targetNumber := flag.Int("targetNumber", 700, "The number we're attempting to find a solution for.")
	targetTolerance := flag.Float64("targetTolerance", 0.0000001, "How close the correct answer should be to count.")
	flag.Parse()

	//reader := bufio.NewReader(os.Stdin)
	population := make([]genome, *popSize)

	// Initial pass to create completely random population
	for i := range population {
		population[i] = make(genome, *genomeSize)
		for j := 0; j < *genomeSize; j++ {
			population[i][j] = uint8(rand.Uint32() & 0xf)
		}
	}

	generationCount := 1

	for {
		// Calculate fitness
		fitnessScores := make([]float64, *popSize)
		fitnessTotal := 0.0
		for i := range population {
			fitnessScores[i] = population[i].Fitness(*targetNumber)
			fitnessTotal += fitnessScores[i]
		}

		// Print out the current generation champion
		championIndex := LargestIndex(fitnessScores)
		fmt.Printf("Generation %d: %s = %f\n", generationCount, population[championIndex].String(), population[championIndex].Value())

		// Check to see if we made it
		if float64(*targetNumber) - *targetTolerance < population[championIndex].Value() && population[championIndex].Value() < float64(*targetNumber) + *targetTolerance {
			fmt.Printf("Solution found!\n")
			os.Exit(0)
		}
		//_, _ = reader.ReadString('\n')

		// begin the next population
		newPopulation := make([]genome, *popSize)
		for i := range newPopulation {
			if i == 0 && *immortalChampion {
				newPopulation[i] = population[championIndex]
			} else {
				newPopulation[i] = Breed(population, fitnessScores, fitnessTotal, *mutRate)
			}
		}

		population = newPopulation
		generationCount++
	}

}
