package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
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
	clock := time.After(5 * time.Second)
	var correct int
	for _, v := range problems {
		fmt.Println("Solve", v[0])
		var c chan string = make(chan string)
		go func() {
			var input string
			fmt.Scanf("%s\n", &input)
			c <- input
		}()

		select {
		case <-clock:
			fmt.Println("End of quiz, too slow")
			fmt.Print("You got ", correct, "/", len(problems))
			os.Exit(0)

		case answer := <-c:
			if answer == v[1] {
				correct += 1
			}
		}
	}
	fmt.Print("You got ", correct, "/", len(problems))
	os.Exit(0)

}
