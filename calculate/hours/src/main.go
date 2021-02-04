// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("RIT_INPUT_1ST_ENTRY")
	input2 := os.Getenv("RIT_INPUT_1ST_EXIT")
	input3 := os.Getenv("RIT_INPUT_2ND_ENTRY")
	input4 := os.Getenv("RIT_INPUT_2ND_EXIT")

	formula.Formula{
		Entry1: input1,
		Exit1:  input2,
		Entry2: input3,
		Exit2:  input4,
	}.Run()
}
