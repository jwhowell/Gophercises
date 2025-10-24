package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	Questions []string
	Answers   []string
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("error reading csv %v", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse csv file: %v", err)
	}
	fmt.Println(lines)
}
