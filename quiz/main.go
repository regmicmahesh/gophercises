package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}

var questions = make([]Question, 0)

func main() {
	var (
		fileName = flag.String("file", "problems.csv", "CSV File to Read Problems From")
		duration = flag.Int64("time", 5, "Time Duration To Wait for Answer")
	)

	flag.Parse()

	f, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("Unable to open file")
		return
	}

	csvReader := csv.NewReader(f)

	for {

		rawText, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		questions = append(questions, Question{Text: rawText[0], Answer: strings.TrimSpace(rawText[1])})
	}
	userAnswer := make(chan string)
	var answer string
	var correctAnswer int
	timer := time.NewTimer(time.Duration(*duration) * time.Second)
	for i, qn := range questions {

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
	fmt.Printf("You answered %d/%d questions correctly\n", correctAnswer, len(questions))
}
