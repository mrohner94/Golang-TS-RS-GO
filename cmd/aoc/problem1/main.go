package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Problem 1 : Submarine navigation

// Without all the fluff and festive fun of advent of code here is th deets.

//     Your submarine starts at position 0,0 (x position, depth)
//     Parse input to direct your submarine.
//     Multiply your depth * x progression to get an answer

// Example input

// forward 5
// down 5
// forward 8
// up 3
// down 8
// forward 2

// Note

// this means that up 3 is actually -3 from the y axis.

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Coordinate struct {
	x int
	y int
}

func parseLine(line string) Coordinate {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal("this should never happen")
	}

	switch parts[0] {
	case "forward":
		return Coordinate{
			x: amount,
			y: 0,
		}
	case "up":
		return Coordinate{
			x: 0,
			y: -amount,
		}
	case "down":
		return Coordinate{
			x: 0,
			y: amount,
		}
	default:
		return Coordinate{
			x: 0,
			y: 0,
		}
	}
}

func main() {
	lines := strings.Split(getInput(), "\n")
	coords := Coordinate{0, 0}

	for _, line := range lines {
		amount := parseLine(line)
		coords.x += amount.x
		coords.y += amount.y
	}

	answer := coords.x * coords.y

	fmt.Printf("point: %+v\nanswer: %d", coords, answer)
}
