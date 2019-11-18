# Go - Genetic Math

A small experiment with genetic algorithms in Go.

Example usage: `go run main.go genome.go --target 200`

![Example](https://i.imgur.com/uibDL26.png)

## The genomes

Each genome is a series of `genomeSize` genes. Each gene is a 4-bit opcode.

- 0 - 9 is a single digit.
- 10 (+) is addition.
- 11 (-) is subtraction.
- 12 (*) is multiplication.
- 13 (/) is division.
- 14 - 15 (X) is an invalid opcode and is ignored.

An example genome is `3+2*4/2`.

Operations are resolved left to right. Duplicate numbers, operations, and invalid opcodes are ignored.

For example:

- `3+2*4/2` = ((3+2) * 4) / 2 = 10
- `32+*1X9` = 3 (ignored) + (ignored) 1 (ignored) (ignored) = 4

## Flags

- `target` float
    - The value the genomes are attempting to solve for. (default 200)
- `populationSize` int
    - The number of genomes in a population. (default 100)
- `mutationRate` float
    - The chance of a gene mutating when breeding. (default 0.01)
- `immortalChampion`
    - If set, the fittest genome in a population will automatically survive to the next. (default true)
- `genomeSize` int
    - The number of genes (operations) in a genome. (default 7)
- `crossover`
    - If set, breeding a new genome will involve two parents with crossover. (default true)
- `colorOutput`
    - If set, genomes will be printed colourized, with invalid genes printed in red. (default true)
- `diffOutput`
    - If set, champions will only be printed if they're different to the previous generation's champion. (default true)
