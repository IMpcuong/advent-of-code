package main

import (
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day9.txt
var day9Data string

type Movement struct {
	Direction string
	Steps     int
}

func convertToMovements(input string) []Movement {
	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(input, "\r\n")
	} else {
		lines = strings.Split(input, "\n")
	}

	var movements = make([]Movement, 0)
	for _, line := range lines {
		direction := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		m := Movement{
			Direction: direction,
			Steps:     steps,
		}
		movements = append(movements, m)
	}
	return movements
}

func getSign(num float64) float64 {
	if num > 1 {
		return 1
	}
	if num < -1 {
		return -1
	}
	return num
}

// NOTE: `Knots := {Head(x, y), Tail(x, y)}`.

type Knot struct {
	x float64
	y float64
}

func (k *Knot) New(newX, newY float64) *Knot {
	return &Knot{
		x: newX,
		y: newY,
	}
}

func (k *Knot) MoveWithDirection(direction string) {
	switch direction {
	case "R": // Right.
		k.x += 1
	case "L": // Left.
		k.x -= 1
	case "U": // Up.
		k.y += 1
	case "D": // Down.
		k.y -= 1
	default:
		panic("Error: Directional opaque!")
	}
}

func (k Knot) DiffEachAxis(other Knot) (float64, float64) {
	return other.x - k.x, other.y - k.y
}

func (k *Knot) TailMoveToward(newHead Knot) {
	tail := *k
	disX, disY := newHead.DiffEachAxis(tail)
	absX, absY := math.Abs(disX), math.Abs(disY)

	// NOTE: If the `newHead` and the `oldTail` (k) have the maximum
	// 	distance (in absolute) between them are between the range of [-1; +1],
	// 	then we won't need to update anything.
	if math.Max(absX, absY) <= 1 {
		return
	}

	// NOTE: In each instruction of motion, the rope is moved by an atomic distance,
	//	or more specifically, by routing along each axis within the boundary of an unit.
	// if absY > absX {
	// 	k.x = newHead.x
	// 	k.y = newHead.y - getSign(disY)
	// } else if absY < absX {
	// 	k.y = newHead.y
	// 	k.x = newHead.x - getSign(disX)
	// } else {
	// 	// NOTE: `disX == disY`.
	// 	k.x = newHead.x - getSign(disX)
	// 	k.y = newHead.y - getSign(disY)
	// }
	k.x += getSign(newHead.x)
	k.y += getSign(newHead.y)
}

func solvingDay9Part1(data string) int {
	motions := convertToMovements(data)
	knots := [2]Knot{} // NOTE: `Rope := { (Head, Tail) := (Knot(x0, y0), Knot(x1, y1)) }`.
	hashSet := make(map[Knot]struct{})

	// Solution2:
	// trace := []Knot{}
	// trace = append(trace, knots[1]) // NOTE: Tracing after each tail's last position, and recorded by an array.

	for _, m := range motions[:20] {
		for step := 0; step < m.Steps; step++ {
			for idx := range knots {
				// NOTE: `head` move normally.
				if idx == 0 {
					knots[idx].MoveWithDirection(m.Direction)
					continue
				}

				head := knots[idx-1]
				knots[idx].TailMoveToward(head)
			} // Move both knots.

			hashSet[knots[len(knots)-1]] = struct{}{} // Adding each knot-tail's position into the hash set.
			fmt.Println(m)
			fmt.Println(knots)
			// fmt.Println(hashSet)
		} // Finish one motion from one instruction.
	} // Finish all motions as the instruction list.

	fmt.Println(knots)

	return len(hashSet)
}
