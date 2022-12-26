package main

import (
	_ "embed"
	"math"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day9_sample.txt
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

func isNegative(num int) bool {
	return math.Signbit(float64(num))
}

// NOTE: `Knots := {Head(x, y), Tail(x, y)}`.

type Knot struct {
	x int
	y int
}

func (k *Knot) New(newX, newY int) *Knot {
	return &Knot{
		x: newX,
		y: newY,
	}
}

func (k *Knot) MoveWithDirection(m Movement) {
	switch m.Direction {
	case "R":
		k.x += m.Steps
	case "L":
		k.x -= m.Steps
	case "U":
		k.y -= m.Steps
	case "D":
		k.y += m.Steps
	default:
		panic("Error: Directional opaque!")
	}
}

func (k Knot) DistanceWith(other Knot) (int, int) {
	return other.x - k.x, other.y - k.y
}

func (k *Knot) AdjustTail(newHead Knot) {
	tail := *k
	disX, disY := newHead.DistanceWith(tail)
	absX, absY := math.Abs(float64(disX)), math.Abs(float64(disY))

	// NOTE: If the `newHead` and the `oldTail` (k) have the maximum
	// distance (in absolute) between them in the range of -/+1, then
	// we won't need to update anything.
	if math.Max(absX, absY) <= 1 {
		return
	}

	if absY > absX {
		k.x = newHead.x
		if isNegative(disY) {
			k.y = newHead.y + disY
		} else {
			k.y = newHead.y - disY
		}
	} else if absY < absX {
		if isNegative(disX) {
			k.x = newHead.x + disX
		} else {
			k.x = newHead.x - disX
		}
		k.y = newHead.y
	} else {
		// FIXME: 4 negative condition waiting to be checked.
		// NOTE: `disX == disY`.
		if isNegative(disX) && isNegative(disY) {
			k.x = newHead.x + disX
			k.y = newHead.y + disY
		} else {
			k.x = newHead.x - disX
			k.y = newHead.y - disY
		}
	}
}
