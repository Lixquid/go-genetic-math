# Go - Genetic Math

A small experiment with genetic algorithms in Go.

Example usage: `go run main.go --targetNumber 30`

```
$> go run main.go --targetNumber 1423 --genomeSize 14
Generation 1: 91*X9X99747485 = 81.000000
Generation 4: 8897244+X7*71- = 105.000000
Generation 23: 88972*47X7*71- = 224.000000
Generation 29: 88972*-7X2*71- = 392.000000
Generation 38: 86972*4*X7*71- = 1568.000000
Generation 41: 86972*4*X7*61- = 1344.000000
Generation 334: 86972*4*X7*6+9 = 1353.000000
Generation 671: 86972*4*97*579 = 1440.000000
Generation 688: 86972*4*97*5-9 = 1431.000000
Generation 25942: 86*52*4*97-5-9 = 1426.000000
Generation 25949: 86*52*4*97-7-9 = 1424.000000
Generation 25951: 86*52*4*97-8-9 = 1423.000000
Solution found!
```

## Flags

- `populationSize`: The number of genomes in each population. Defaults to `100`.
- `mutationRate`: The mutation rate for each gene. When breeding, this is the chance each gene will mutate into a new gene. Defaults to `0.01`.
- `immortalChampion`: If set, the genome with the highest fitness will automatically survive into the next generation. Defaults to on.
- `genomeSize`: The number of genes (instructions) in each genome. Defaults to `7`.
- `targetNumber`: The number the genomes are attempting to find a solution for. Defaults to `200`.

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
