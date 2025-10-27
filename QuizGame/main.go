package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	Questions []string
	Answers   []string
}

func main() {
	var filename string
	flag.StringVar(&filename, "Filename", "problems.csv", "Filepath to quiz csv, format Q,A")

	flag.Parse()

	file, err := os.Open(filename)
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
		fmt.Printf("\nQuestion Number #%d: %v ", n+1, l[0])
		fmt.Scanf("%s\n", &input)
		if input == l[1] {
			fmt.Println("Correct")
		} else {
			fmt.Printf("Incorrect, the answer is: %v\n", l[1])
		}
	}
}
