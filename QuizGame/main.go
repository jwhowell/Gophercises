package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Questions []string
	Answers   []string
}

func main() {
	filename := flag.String("Filename", "problems.csv", "Filepath to quiz csv, format Q,A")
	timeLimit := flag.Int("Limit", 30, "Enter a time limit (seconds) for the quiz, default 30 seconds")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("error reading csv %v", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse csv file: %v", err)
	}
	fmt.Print(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var correct int
	for n, l := range lines {
		fmt.Printf("\nQuestion Number #%d: %v ", n+1, l[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("%d correct out of %d", correct, len(lines))
			return
		case answer := <-answerCh:
			if answer == l[1] {
				fmt.Println("Correct")
				correct++
			} else {
				fmt.Printf("Incorrect, the answer is: %v\n", l[1])
			}
		}
	}

}
