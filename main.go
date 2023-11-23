package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

// checkCodeCoverage runs the `go tool cover --func cover.out` command
// to extract the total code coverage percentage. It compares the coverage
// with the desired coverage threshold and prints a success or failure message
// accordingly. The function takes an integer parameter `desiredCoverage`
// which represents the desired coverage threshold.
func checkCodeCoverage(desiredCoverage int) {
	// Run the `go tool cover --func cover.out` command and capture the output
	out, err := exec.Command("go", "tool", "cover", "--func", "cover.out").Output()
	if err != nil {
		// Print the error message to stderr and exit with a non-zero status code
		fmt.Fprint(os.Stderr, "Error: ", string(err.(*exec.ExitError).Stderr))
		os.Exit(1)
	}

	// Define a regular expression to extract the total coverage percentage
	re := regexp.MustCompile(`total:.+\W+(\d+\.\d+)%`)

	// Extract the total coverage from the output using the regular expression
	totalCoverage, err := strconv.ParseFloat(re.FindStringSubmatch(string(out))[1], 64)
	if err != nil {
		// Panic if there was an error parsing the coverage percentage
		panic(err)
	}

	if totalCoverage < float64(desiredCoverage) {
		// Print a failure message if the coverage is less than the desired threshold
		fmt.Printf(
			"❌ Code coverage is %.1f%% but needs to be %d%% or more to pass\n",
			totalCoverage, desiredCoverage,
		)
		os.Exit(1)
	}

	// Print a success message if the coverage is equal to or higher than the
	// desired threshold
	fmt.Printf("✅ Code coverage is %.1f%%\n", totalCoverage)
}

func main() {
	// Set the desired code coverage threshold
	desiredCoverage := 80

	// Check the code coverage
	checkCodeCoverage(desiredCoverage)
}
