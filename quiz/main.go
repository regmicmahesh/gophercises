package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type question struct {
	Text   string
	Answer string
}

func parseCSV(fileName string) *[]question {

	f, _ := os.Open(fileName)
	rdr := csv.NewReader(f)

	allLines, _ := rdr.ReadAll()

	questions := make([]question, len(allLines))

	for i, v := range allLines {
		questions[i] = question{Text: v[0], Answer: strings.TrimSpace(v[1])}
	}

	return &questions

}

func main() {

	var (
		fileName = flag.String("file", "problems.csv", "CSV File to Read Problems From")
		duration = flag.Int64("time", 5, "Time Duration To Wait for Answer")
	)

	flag.Parse()

	userAnswer := make(chan string)
	var answer string
	var correctAnswer int
	questions := parseCSV(*fileName)

	// this makes a timer which sends me a message in timer.C, after the time
	// is passed.
	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	for i, qn := range *questions {
		fmt.Printf("Question #%d: %s", i+1, qn.Text)

		go func() {
			fmt.Printf("\nYour Answer: ")
			fmt.Scanf("%s", &answer)
			userAnswer <- answer
		}()

		select {

		case <-timer.C:
			fmt.Println("\nYour time has expired")
			os.Exit(0)

		case answer = <-userAnswer:
			if answer == qn.Answer {
				correctAnswer++
			}
		}

	}
	fmt.Printf("You answered %d/%d questions correctly\n", correctAnswer, len(*questions))
}
