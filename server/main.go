package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"solver/dpll"
	"solver/formula"
)

const PROMPT string = "Clause per line, enter empty line to start solving\n"

func runDPLL(input [][]string) {
	f := formula.NewFormula(input)
	fmt.Printf("Form: %s\n\n", f.String())
	valid, assignments, dc := dpll.Solve(f)
	var dcString []string
	for s := range dc {
		dcString = append(dcString, s)
	}

	if valid {
		var trues []string
		var falses []string
		for str, assignedValue := range assignments {
			if assignedValue {
				trues = append(trues, str)
			} else {
				falses = append(falses, str)
			}
		}
		fmt.Printf("True: %v\nFalse: %v\nDon't Care: %v\n", trues, falses, dcString)
	} else {
		fmt.Printf("unsat")
	}
}

func scanInput() [][]string {
	var lines [][]string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(PROMPT)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		// empty newline or ctrl-D
		if input == "" || input == "q" || scanner.Err() != nil {
			break
		}
		re := regexp.MustCompile(`\s+`)
		literalsList := re.Split(input, -1)
		lines = append(lines, literalsList)
	}

	return lines
}

func main() {
	input := scanInput()
	runDPLL(input)
	bufio.NewScanner(os.Stdin).Scan() // hold window if using exe
}
