package main

import (
	_ "embed"
	"math"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day09.txt
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

func getEdgeVal(num, min, max float64) float64 {
	if num > max {
		num = max
	}
	if num < min {
		num = min
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
	return k.x - other.x, k.y - other.y
}

// Idea from: https://github.com/sluongng/advent2022/blob/31c1b99ec28432aac144adfad50c7b9870a95888/day9/src/main.rs#L37

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
	if absY > absX {
		k.x = newHead.x
		k.y = newHead.y - getEdgeVal(disY, -1, 1)
	} else if absY < absX {
		k.y = newHead.y
		k.x = newHead.x - getEdgeVal(disX, -1, 1)
	} else {
		// NOTE: `absX == absY`.
		k.x = newHead.x - getEdgeVal(disX, -1, 1)
		k.y = newHead.y - getEdgeVal(disY, -1, 1)
	}

	// Solution2: Inspired from https://github.com/sluongng/advent2022/blob/fe20a9bbe3bcbe4580f54dde6e71429c3984e807/day9/src/main.rs#L48

	// k.x += getEdgeVal(disX, -1, 1)
	// k.y += getEdgeVal(disY, -1, 1)

	// End From: https://github.com/sluongng/advent2022/blob/fe20a9bbe3bcbe4580f54dde6e71429c3984e807/day9/src/main.rs#L49
}

// End From: https://github.com/sluongng/advent2022/blob/31c1b99ec28432aac144adfad50c7b9870a95888/day9/src/main.rs#L51

func solvingDay9(data string, ropeLen int) int {
	motions := convertToMovements(data)
	knots := make([]Knot, ropeLen) // NOTE: `Rope := { (Head, Tail) := (Knot(x0, y0), Knot(xN, yN)) | 0 <= i <= N=ropeLen }`.
	hashSet := make(map[Knot]struct{})

	// Solution2:
	// trace := []Knot{}
	// trace = append(trace, knots[len(knots)-1]) // NOTE: Tracing after each tail's last position, and recorded by an array.

	for _, m := range motions[:] {
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
		} // Finish one motion from one instruction.
	} // Finish all motions as the instruction list was told.

	return len(hashSet)
}
