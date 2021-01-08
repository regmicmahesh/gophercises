package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Question struct {
	Text   string
	Answer string
}

type Questions []Question

var questions Questions

func main() {
	var (
		fileName = flag.String("file", "problems.csv", "CSV File to Read Problems From")
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

	var answer string
	var correctAnswer int
	for i, qn := range questions {
		fmt.Printf("Question #%d: %s", i+1, qn.Text)
		fmt.Printf("\nYour Answer: ")
		fmt.Scanf("%s", &answer)
		if answer == qn.Answer {
			correctAnswer++
		}
	}
	fmt.Printf("You answered %d/%d questions correctly\n", correctAnswer, len(questions))
}
