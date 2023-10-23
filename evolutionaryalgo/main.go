package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	targetString   = "1101000011"
	populationSize = 100
	mutationRate   = 0.01
)

func main() {
	// set random generator
	rand.Seed(time.Now().UnixNano())

	// Initialize the population with random binary strings
	population := initializePopulation(populationSize, len(targetString))

	fmt.Printf("init population ,100 individus --> %s \n", population)
	fmt.Printf("target individu --> %s \n", targetString)

	generations := 0
	for {
		// Evaluate the fitness of each individual in the population
		fitnessScores := calculateFitness(population)

		// Check for a solution
		bestIndividual, bestFitness := findBestIndividual(population, fitnessScores)
		if bestFitness == len(targetString) {
			fmt.Printf("Found a solution in %d generations: %s\n", generations, bestIndividual)
			break
		}

		fmt.Printf("generation %d best individual -->%s\n ", generations, bestIndividual)

		// Create a new population through selection and reproduction
		newPopulation := make([]string, populationSize)
		for i := range newPopulation {
			parent1 := selectIndividual(population, fitnessScores)
			parent2 := selectIndividual(population, fitnessScores)
			child := crossover(parent1, parent2)
			child = mutate(child, mutationRate)
			newPopulation[i] = child
		}
		population = newPopulation

		generations++
	}
}

func initializePopulation(size, length int) []string {
	population := make([]string, size)
	for i := range population {
		population[i] = randomBinaryString(length)
	}
	return population
}

func randomBinaryString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		if rand.Float64() < 0.5 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}

func calculateFitness(population []string) []int {
	fitnessScores := make([]int, len(population))
	for i, individual := range population {
		for j := range individual {
			if individual[j] == targetString[j] {
				fitnessScores[i]++
			}
		}
	}
	return fitnessScores
}

func findBestIndividual(population []string, fitnessScores []int) (string, int) {
	bestIndex := 0
	bestFitness := fitnessScores[0]
	for i, fitness := range fitnessScores {
		if fitness > bestFitness {
			bestIndex = i
			bestFitness = fitness
		}
	}
	return population[bestIndex], bestFitness
}

func selectIndividual(population []string, fitnessScores []int) string {
	totalFitness := 0
	for _, fitness := range fitnessScores {
		totalFitness += fitness
	}

	threshold := rand.Intn(totalFitness)
	cumulativeFitness := 0
	for i, fitness := range fitnessScores {
		cumulativeFitness += fitness
		if cumulativeFitness >= threshold {
			return population[i]
		}
	}

	// Fallback if nothing is selected (shouldn't happen)
	return population[0]
}

func crossover(parent1, parent2 string) string {
	// Perform one-point crossover
	crossoverPoint := rand.Intn(len(parent1))
	child := parent1[:crossoverPoint] + parent2[crossoverPoint:]
	return child
}

func mutate(individual string, mutationRate float64) string {
	mutated := []byte(individual)
	for i := range mutated {
		if rand.Float64() < mutationRate {
			if mutated[i] == '0' {
				mutated[i] = '1'
			} else {
				mutated[i] = '0'
			}
		}
	}
	return string(mutated)
}
