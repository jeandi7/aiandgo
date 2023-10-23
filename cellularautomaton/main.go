package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Width  = 8
	Height = 8
)

type Grid [Width][Height]bool

func (g Grid) Print(gen int) {
	fmt.Printf("generation n°(%d)-------------------------\n", gen)

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if g[x][y] {
				fmt.Print("◼")
			} else {
				fmt.Print("◻")
			}
		}
		fmt.Println()
	}

}

func (g Grid) NextGeneration() Grid {
	var newGrid Grid
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			liveNeighbors := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < Width && ny >= 0 && ny < Height && g[nx][ny] {
						liveNeighbors++
					}
				}
			}
			if g[x][y] {
				newGrid[x][y] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				newGrid[x][y] = liveNeighbors == 3
			}
		}
	}
	return newGrid
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var grid Grid
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			grid[x][y] = rand.Intn(2) == 1
		}
	}
	for gen := 0; gen < 10; gen++ {
		grid.Print(gen)
		grid = grid.NextGeneration()

	}
}
