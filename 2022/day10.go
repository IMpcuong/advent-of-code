package main

import (
	_ "embed"
	"fmt"
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

func solvingDay10P1(input string, divisor int) (int, int) {
	var valInReg int = 1
	var cyclePos int

	result := 0
	signals := convertToSignal(input)
	for _, s := range signals {
		// fmt.Printf("%#v\t%#v\n", s.Cycles, *(s.Instruction))

		cyclePos += s.Cycles
		valInReg += (*s.Instruction).Value
		fmt.Println(cyclePos, valInReg)

		if cyclePos%divisor == 20 {
			val := cyclePos * valInReg
			fmt.Println("=====>\t", cyclePos, result)
			result += val
		}
	}

	return result, cyclePos
}
