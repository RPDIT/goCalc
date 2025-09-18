package main

import (
	"fmt"
	"log"
	"strconv"

	L "example.com/calculator_app/lib"
)

type Application struct {
	input  []string
	status bool
}

func (a Application) appendInput(str string) []string {
	return append(a.input, str)
}

func (a Application) resetInput() []string {
	return make([]string, 0)

}

func (a Application) getInput() Application {
	input_status := true
	var counter int = 0
	
	for input_status {
		var local_input string
		counter = counter + 1
		
		if counter > 3 {
			input_status = false
			break
		}
		
		_, err := fmt.Scan(&local_input)
		if err != nil {
			log.Printf("Failed to parse input: %v", err)
		}

		a.input = a.appendInput(local_input)
	}
	return a
}


func main() {
	a := Application{status: true}
	fmt.Print("Please enter your values separated by the command. \nExample: \n  7 * 7\n")
	for a.status {
		a.input = a.resetInput()
		a = a.getInput()
		test_result := L.FunctionChooser(a.input[1], ParseFloatsfromStrings(a.input[0]), ParseFloatsfromStrings(a.input[2]))
		if test_result.Err != nil {
			log.Printf("Malformed input: %v\n", test_result.Err)
		}
		fmt.Printf("Got this value: %v\n", test_result.Value)
	}

}

func ParseFloatsfromStrings(value string) float64 {
	if s, err := strconv.ParseFloat(value, 32); err == nil {
		return float64(s)
	}
	return 0
}
