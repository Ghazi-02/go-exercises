package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readFile(f string) [][]string {
	file, err := os.Open(f)
	defer file.Close()

	readVal := csv.NewReader(file)
	recordedVals, err := readVal.ReadAll()
	if err != nil {
		fmt.Println(err)

	}

	return recordedVals
}

func main() {
	args := os.Args
	// fmt.Println(args[1:])
	var quizFile string
	if len(args) == 2 {
		quizFile = args[1]
	} else {
		quizFile = "problems.csv"
	}
	problems := readFile(quizFile)
	var correct int
	for i := range problems {

		
		fmt.Println("Solve",problems[i][0])
		var input string
		fmt.Scanf("%s\n", &input)
		if input == problems[i][1] {
			correct += 1
		}

	}
	fmt.Print("You got ",correct, "/", len(problems))
	os.Exit(0)

}
