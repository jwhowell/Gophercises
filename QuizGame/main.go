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
	fmt.Print(lines)
	var input string
	for n, l := range lines {
		fmt.Printf("\nQuestion Number #%d: %v ", n, l[0])
		fmt.Scanf("%s\n", &input)
		if input == l[1] {
			fmt.Println("Correct")
		} else {
			fmt.Printf("Incorrect, the answer is: %v\n", l[1])
		}
	}
}
