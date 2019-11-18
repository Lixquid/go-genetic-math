package main

import (
	"math"
	"math/rand"
	"strings"
)

type Genome []uint8

const (
	EscapeRed   = "\x1b[31m"
	EscapeClear = "\x1b[0m"
)

func (g Genome) String(color bool) string {
	var output strings.Builder
	expectDigit := true

	for i := range g {
		switch g[i] {
		case 10:
			if color {
				if expectDigit {
					output.WriteString(EscapeRed)
				} else {
					output.WriteString(EscapeClear)
				}
				expectDigit = true
			}
			output.WriteRune('+')
		case 11:
			if color {
				if expectDigit {
					output.WriteString(EscapeRed)
				} else {
					output.WriteString(EscapeClear)
				}
				expectDigit = true
			}
			output.WriteRune('-')
		case 12:
			if color {
				if expectDigit {
					output.WriteString(EscapeRed)
				} else {
					output.WriteString(EscapeClear)
				}
				expectDigit = true
			}
			output.WriteRune('*')
		case 13:
			if color {
				if expectDigit {
					output.WriteString(EscapeRed)
				} else {
					output.WriteString(EscapeClear)
				}
				expectDigit = true
			}
			output.WriteRune('/')
		case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			if color {
				if expectDigit {
					output.WriteString(EscapeClear)
				} else {
					output.WriteString(EscapeRed)
				}
				expectDigit = false
			}
			output.WriteRune(rune(g[i] + 48))
		default:
			// Invalid gene
			if color {
				output.WriteString(EscapeRed)
			}
			output.WriteRune('X')
		}
	}

	output.WriteString(EscapeClear)

	return output.String()
}

const (
	none     = iota
	add      = iota
	subtract = iota
	multiply = iota
	divide   = iota
)

func (g Genome) Value() float64 {
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

func (g Genome) Fitness(target float64) float64 {
	epsilon := math.Nextafter(1.0, 2.0) - 1.0
	return 1 / math.Max(math.Abs(target-g.Value()), epsilon)
}

func (g Genome) Equal(other Genome) bool {
	for i := range g {
		if other[i] != g[i] {
			return false
		}
	}
	return true
}

func selectGenome(population []Genome, fitness []float64, fitnessTotal float64) Genome {
	fitnessGoal := rand.Float64() * fitnessTotal
	total := 0.0
	for i := range population {
		if total+fitness[i] < fitnessGoal {
			total += fitness[i]
			continue
		}

		return population[i]
	}

	// Should never ever happen: means sum of population fitness != fitnessTotal
	panic("Reached end of population when selecting Genome")
}

func Breed(population []Genome, fitness []float64, fitnessTotal float64, mutationRate float64) Genome {
	new := make(Genome, len(population[0]))
	copy(new, selectGenome(population, fitness, fitnessTotal))

	// TODO: Crossover

	// Mutate - we're actually randomizing an entire gene, instead of mutating one bit
	for j := range new {
		if rand.Float64() < mutationRate {
			new[j] = uint8(rand.Uint32() & 0xf)
		}
	}

	return new
}
