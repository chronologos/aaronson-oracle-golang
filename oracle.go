package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	oracle()

}

const DEBUG = false

var correct, total uint32
var guess string

func stringify(list []string) string {
	var result string
	for _, value := range list {
		result += value
	}
	return result
}

func randGuess() string {
	random := rand.Int() % 2
	if random == 1 {
		guess = "f"
		fmt.Println("ORACLE guesses f")
	} else {
		guess = "g"
		fmt.Println("ORACLE guesses: g")
	}
	return guess
}

func oracle() {
	var history []string // we will use this as queue of length 5
	oracle_map := make(map[string]string)
  fmt.Println("~~~WELCOME TO THE ORACLE~~~")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("---------\nf or g?: ")
		text, _ := reader.ReadString('\n')
		text = string(text[0])

		if text == "p" {
			fmt.Printf("{")
			for key, value := range oracle_map {
				fmt.Printf("%q: %q ,", key, value)
			}
			fmt.Printf("}")
			continue
		}

		if len(history) > 5 {
			history = history[1:] //discard top element
      if DEBUG {
        fmt.Printf("%q", history)
      }

			if value, ok := oracle_map[stringify(history)]; !ok {
				guess = randGuess()
				if text == guess {
					correct += 1
				}
			} else {

				guess = value
				fmt.Printf("The oracle guesses: %s", guess)
				if text == guess {
					correct += 1
				}
			}
			oracle_map[stringify(history)] = text

		} else {
			guess = randGuess()
      if text == guess {
        correct += 1
      }
		}
    total++

		fmt.Printf("Correct %d, Total %d Percentage right = %f \n", correct, total, float64(correct)/float64(total))
		history = append(history, text) // why initialize again?
	}
	return
}
