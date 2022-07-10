package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type pair struct {
	question string
	answer   string
}
type result struct {
	questions int
	correct   int
}

func main() {
	file, timer := parseFlags()
	quiz := readCsv(file)
	result := launchQuiz(quiz, timer)
	showresult(result)
}

func showresult(r result) {
	fmt.Println("Total questions:", r.questions)
	fmt.Println("Correct answers:", r.correct)
}

func launchQuiz(quiz []pair, timer int) result {
	fmt.Println("Are you ready kids? (y,n)")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := strings.ToLower(strings.Trim(scanner.Text(), " "))
		if text == "y" {
			break
		} else {
			fmt.Println("I can't hear you!")
		}
	}
	var result result
	result.questions = len(quiz)
	for _, pair := range quiz {
		fmt.Println(pair.question)
		scanner.Scan()
		answer := strings.Trim(scanner.Text(), " ")
		if answer == pair.answer {
			result.correct++
		}
	}
	return result
}

func readCsv(file string) []pair {
	quizfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	var quiz []pair
	reader := csv.NewReader(quizfile)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		q := strings.Join(append([]string{}, strings.Join(rec[0:(len(rec)-1)], ",")), ", ")
		a := rec[len(rec)-1]
		quiz = append(quiz, pair{question: q, answer: a})
	}
	return quiz
}

func parseFlags() (string, int) {
	file := flag.String("f", "problems.csv", "path to a file with quiz problem")
	timer := flag.Int("t", 30, "timer")
	flag.Parse()
	return *file, *timer
}
