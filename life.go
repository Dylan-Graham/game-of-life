package main

import (
	"fmt"
	"math/rand"
	"time"
)

func displayWorld(world [20][20]bool) {
	for i := 1; i < len(world)-1; i++ {
		for j := 1; j < len(world)-1; j++ {
			if world[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func runGeneration(world [20][20]bool) [20][20]bool {
	/*
	* for each cell in the world run the 3 rules:
	*
	* 1) any live cell with 2/3 neighbours survives
	* 2) any dead cell with 3 live neighbours becomes a live cell
	* 3) all other live cells die in the next generation (similarly, all dead cells stay dead)
	*
	 */

	var temp = world

	for i := 1; i < len(world)-1; i++ {
		for j := 1; j < len(world)-1; j++ {

			neighbours := 0

			if world[i-1][j-1] {
				neighbours++
			}
			if world[i-1][j] {
				neighbours++
			}
			if world[i-1][j+1] {
				neighbours++
			}

			if world[i][j-1] {
				neighbours++
			}
			if world[i][j+1] {
				neighbours++
			}

			if world[i+1][j-1] {
				neighbours++
			}
			if world[i+1][j] {
				neighbours++
			}
			if world[i+1][j+1] {
				neighbours++
			}

			// Life/Death logic
			if world[i][j] {
				if neighbours != 2 && neighbours != 3 {
					temp[i][j] = false
				} else {
					temp[i][j] = true
				}
			} else {
				if neighbours == 3 {
					temp[i][j] = true
				} else {
					temp[i][j] = false
				}
			}

			// if live cell & 2/3 neighbours then survive
			// else die

			// if dead cell & 3 neighbours then come alive
			// else stay dead

		}
	}

	return temp
}

func initWorld(world [20][20]bool) [20][20]bool {
	for i := 1; i < len(world)-1; i++ {
		for j := 1; j < len(world)-1; j++ {
			rand.Seed(time.Now().UnixNano())
			world[i][j] = rand.Intn(2) == 0
		}
	}
	return world
}

func main() {
	/* an array with 5 rows and 5 columns*/
	var world [20][20]bool

	world = initWorld(world)
	lifeCycles := 200

	for i := 0; i < lifeCycles; i++ {
		fmt.Print("\x0c") // Clear screen
		displayWorld(world)
		world = runGeneration(world)
		time.Sleep(time.Second / 30)
	}
}
