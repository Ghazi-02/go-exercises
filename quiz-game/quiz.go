package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readFile() [][]string {
	file, err := os.Open("problems.csv")
	defer file.Close()

	
	readVal := csv.NewReader(file)
	recordedVals,err := readVal.ReadAll()
	if err != nil {
		fmt.Println(err)

	}
	
	return recordedVals
}
func solver()

func main() {
	problems := readFile()
	fmt.Println(problems[0][1])
	
}