package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

type point struct {
	x int
	y int
	direction int
}

func (p *point) move (direction string) {

	//direction controls where we are currently facing, 0 being north, 3 being west
	if direction[0] == 'L' {
		p.direction--
	} else {
		p.direction++
	}

	//make sure that direction is positive between 0 and 3
	p.direction %= 4
	if p.direction < 0 {
		p.direction = 4 + p.direction
	}

	distance, _ := strconv.Atoi(direction[1:])
	switch p.direction {
		case 0:
			p.y += distance
			break
		case 1:
			p.x += distance
			break
		case 2:
			p.y -= distance
			break
		case 3:
			p.x -= distance
			break

	}
}

func (p *point) distance () int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func main() {
	input := "L5, R1, L5, L1, R5, R1, R1, L4, L1, L3, R2, R4, L4, L1, L1, R2, R4, R3, L1, R4, L4, L5, L4, R4, L5, R1, R5, L2, R1, R3, L2, L4, L4, R1, L192, R5, R1, R4, L5, L4, R5, L1, L1, R48, R5, R5, L2, R4, R4, R1, R3, L1, L4, L5, R1, L4, L2, L5, R5, L2, R74, R4, L1, R188, R5, L4, L2, R5, R2, L4, R4, R3, R3, R2, R1, L3, L2, L5, L5, L2, L1, R1, R5, R4, L3, R5, L1, L3, R4, L1, L3, L2, R1, R3, R2, R5, L3, L1, L1, R5, L4, L5, R5, R2, L5, R2, L1, L5, L3, L5, L5, L1, R1, L4, L3, L1, R2, R5, L1, L3, R4, R5, L4, L1, R5, L1, R5, R5, R5, R2, R1, R2, L5, L5, L5, R4, L5, L4, L4, R5, L2, R1, R5, L1, L5, R4, L3, R4, L2, R3, R3, R3, L2, L2, L2, L1, L4, R3, L4, L2, R2, R5, L1, R2"
	directions := strings.Split(input, ", ")

	position := point {
		x: 0,
		y: 0,
		direction: 0,
	}

	for _, direction := range directions {
		position.move(direction)
	}
	fmt.Println(position.distance())
}
