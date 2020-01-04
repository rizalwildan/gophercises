package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFile))
	}
	reader := csv.NewReader(file)
	line, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file."))
	}

	problem := parseLines(line)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	problemLoop:
	for i, p := range problem {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <-answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("\n You scored %d out of %d.\n", correct, len(problem))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return  ret
}

type problem struct {
	question string
	answer string
}



func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
