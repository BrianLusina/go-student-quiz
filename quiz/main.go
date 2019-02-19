package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	// using a pointer to the actual file, due to Flag parsing a string
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprint("Failed to open the CSV file: %s\n", *csvFileName))
	}

	// read the file with the csv
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse CSV file")
	}

	// parse the problems
	problems := parseLines(lines)

	fmt.Println(problems)

	correct_answers := 0

	// print out problems to user
	for i, problem := range problems  {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string

		fmt.Scanf("%s\n", &answer)

		if answer == problem.answer {
			correct_answers ++
		}
	}

	fmt.Printf("You got %d out of %d correct.\n", correct_answers, len(problems))
}

// Parse lines in the csv and returns the struct with the problems
func parseLines(lines [][] string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// struct that defines the structure of a problem from the CSV file
type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
