package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	out, err := exec.Command(
		"go", "tool", "cover", "--func", "cover.out",
	).Output()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: ", string(err.(*exec.ExitError).Stderr))
		os.Exit(1)
	}

	re := regexp.MustCompile(`total:.+\W+(\d+\.\d+)%`)
	totalCoverage, err := strconv.ParseFloat(
		re.FindStringSubmatch(string(out))[1],
		64,
	)
	if err != nil {
		panic(err)
	}

	if totalCoverage < 80 {
		fmt.Printf(
			"❌ Code coverage is %.1f%%, but needs to be 80%% or more to pass",
			totalCoverage,
		)
		os.Exit(1)
	} else {
		fmt.Printf("✅ Code coverage is %.1f%%", totalCoverage)
	}
}
