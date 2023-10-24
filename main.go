package main

import (
	"fmt"
	"os"
)

func main() {
	var count int
	var sum float64

	var val float64
	for {

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		count++
	}

	if count > 0 {
		fmt.Println("The avarage is:", sum/float64(count))
	} else {
		fmt.Println("No values were given!")
	}
}
