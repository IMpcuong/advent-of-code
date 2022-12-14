package main

import (
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day10.txt
var day10Data string

type Signal struct {
	Cycles      int     // Maximum looping is required to correctly execute one signal.
	Instruction *RegIns // Simple instructions to guide the register's updating after each cycle.
}

type RegIns struct {
	Type  string // Type := { noop, addx | noop := `do nothing`, addx := `add value to register X` }.
	Value int    // The underlying number that must be added to the X register.
}

func convertToSignal(input string) []Signal {
	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(input, "\r\n")
	} else {
		lines = strings.Split(input, "\n")
	}

	var signals = make([]Signal, 0)
	for _, line := range lines {
		strIns := strings.Split(line, " ")

		insType := strIns[0]
		v := 0
		if len(strIns) > 1 {
			v, _ = strconv.Atoi(strIns[1])
		}
		regIns := &RegIns{
			Type:  insType,
			Value: v,
		}
		var cycles int = len(strIns) // Because the `noop` instruction always stand alone.
		s := Signal{
			Cycles:      cycles,
			Instruction: regIns,
		}

		signals = append(signals, s)
	}

	return signals
}

func signalStrength(cyclePos, regVal, divisor int) int {
	var res int
	if cyclePos%divisor == 20 {
		res = cyclePos * regVal
	} else {
		res = 0
	}
	return res
}

func solvingDay10P1(input string, divisor int) (int, int) {
	var regX int = 1
	var cyclePos int = 1

	result := 0
	signals := convertToSignal(input)
	for _, s := range signals {
		if s.Cycles == 2 {
			result += signalStrength(cyclePos, regX, divisor)
			cyclePos += s.Cycles / 2
			result += signalStrength(cyclePos, regX, divisor)
			cyclePos += s.Cycles / 2
			regX += (*s.Instruction).Value
		} else {
			result += signalStrength(cyclePos, regX, divisor)
			cyclePos += s.Cycles
		}
	}

	return result, cyclePos
}

const (
	COLS int = 40
	ROWS int = 6
)

func renderPixel(screen *[ROWS][COLS]string, cyclePos, regVal int) {
	// NOTE: `1 sprite := 3 pixels`.
	// NOTE: `(cyclePos := [1; 6]) == (COLS := [0; 5])` -> `COL := cyclePos - 1`.
	pixelX := (cyclePos - 1) % COLS // Horizontal axis -> `pixelX := [0; 39]`.
	pixelY := (cyclePos - 1) / COLS // Vertical axis -> `pixelY := COL * [0; 39]`.

	// NOTE: `regVal` == current position of the sprite, counting from the left most pixel.
	if math.Abs(float64(pixelX-regVal)) < 2 {
		(*screen)[pixelY][pixelX] = "#"
	} else {
		(*screen)[pixelY][pixelX] = "."
	}
}

func solvingDay10P2(input string) {
	var regX int = 1
	var cyclePos int = 1

	screen := [ROWS][COLS]string{}
	signals := convertToSignal(input)
	for _, s := range signals {
		if s.Cycles == 2 {
			renderPixel(&screen, cyclePos, regX)
			cyclePos += s.Cycles / 2
			renderPixel(&screen, cyclePos, regX)
			cyclePos += s.Cycles / 2
			regX += (*s.Instruction).Value
		} else {
			renderPixel(&screen, cyclePos, regX)
			cyclePos += s.Cycles
		}
	}

	for row := 0; row < ROWS; row++ {
		fmt.Println(screen[row])
	}
}
