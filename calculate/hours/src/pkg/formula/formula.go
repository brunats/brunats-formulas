// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"time"
)

type Formula struct {
	Entry1 string
	Exit1  string
	Entry2 string
	Exit2  string
}

func (f Formula) Run() {
	firstPeriod, secondPeriod := f.calculatePeriods()

	result := fmt.Sprintf("Hello\n\n")

	result += fmt.Sprintf("Its first calculated period was: %v.\n", firstPeriod.String())

	result += fmt.Sprintf("Its second calculated period was: %v.\n\n", secondPeriod.String())

	result += fmt.Sprintf("Total period was: %v.\n", (firstPeriod + secondPeriod).String())

	fmt.Println(result)
}

func (f Formula) calculatePeriods() (time.Duration, time.Duration) {
	durationEntry1, _ := time.ParseDuration(f.Entry1)
	durationExit1, _ := time.ParseDuration(f.Exit1)
	durationEntry2, _ := time.ParseDuration(f.Entry2)
	durationExit2, _ := time.ParseDuration(f.Exit2)

	firstPeriod := durationExit1 - durationEntry1
	secondPeriod := durationExit2 - durationEntry2

	return firstPeriod, secondPeriod
}
