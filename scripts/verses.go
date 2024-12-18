package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"flag"
)

func main() {
	inputFileName := flag.String("input", "input.md", "input filename")
	outputFileName := flag.String("ouput", "output.json", "output filename")

	flag.Parse()

	fmt.Printf("reading file: %s\n", *inputFileName)

	data, err := ioutil.ReadFile(*inputFileName)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	versesSource := strings.Split(string(data), "\n\n")
	verses := []string{}

	for _, s := range versesSource {
		verses = append(verses, strings.TrimSpace(s))
	}

	json, err := json.MarshalIndent(verses, "", " ")
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}

	err = ioutil.WriteFile(*outputFileName, json, 0644)
	if err != nil {
		log.Fatalf("failed to write json: %v", err)
	}

	fmt.Printf("wrote file: %s\n", *outputFileName)
}
