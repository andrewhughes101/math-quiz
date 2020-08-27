package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, timerlength := readFlags()

	questions, answers := getQuestions(file)

	quiztimer := time.NewTimer(time.Duration(timerlength) * time.Second)

	var score, total int = 0, len(questions)

	answerchn := make(chan string)

	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question: %s\n", questions[i])
		go getAnswer(answerchn)
		timeup := checkAnswer(answerchn, quiztimer, answers[i], &score)
		if timeup {
			break
		}
	}

	fmt.Printf(endGame(score, total))
}

func readFlags() (string, int) {
	filePtr := flag.String("file", "problems.csv", "csv file containing question sand answers")
	timerPtr := flag.Int("timer", 30, "Number of seconds to answer all the questions in")
	flag.Parse()

	return *filePtr, *timerPtr
}

func getQuestions(fileName string) ([]string, []string) {
	questions := make([]string, 0)
	answers := make([]string, 0)

	csvfile, err := os.Open(fileName)
	if err != nil {
		errMsg := fmt.Sprintf("Could not open file %s \n", fileName)
		log.Fatal(errMsg, err)
	}

	reader := csv.NewReader(csvfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}
	return questions, answers
}

func getAnswer(answerchn chan string) {
	var guess string
	fmt.Scanf("%s\n", &guess)
	answerchn <- guess
}

func checkAnswer(answerchn chan string, quiztimer *time.Timer, answer string, score *int) bool {
	select {
	case <-quiztimer.C:
		fmt.Printf("Times up!\n")
		return true
	case guess := <-answerchn:
		if guess == answer {
			fmt.Printf("Correct\n")
			*score++
		} else {
			fmt.Printf("Incorrect, answer is %s\n", answer)
		}
		return false
	}
}

func endGame(score int, total int) string {
	return fmt.Sprintf("You scored %d out of %d\n", score, total)
}
