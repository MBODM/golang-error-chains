package main

import (
	"errors"
	"fmt"
	"os"

	"example.com/error-chains/apperror"
	"example.com/error-chains/package3"
	"example.com/error-chains/simulation"
)

func main() {
	simulationSettings()
	result, err := package3.Execute()
	if err != nil {
		// We differ the errors here, because we wanna do some special printing for our own custom error:
		var customError *apperror.CustomError
		if errors.As(err, &customError) {
			fmt.Println("[Foo] Application Error:", customError)
			os.Exit(1)
		} else {
			fmt.Println("[Bar] Unexpected error(s) occurred:", err)
			os.Exit(1)
		}
	}
	fmt.Println(result)
	fmt.Println()
	fmt.Println("Have a nice day.")
	os.Exit(1)
}

func simulationSettings() {
	// Play around with these settings, to see how things go.
	simulation.NoError = false
	simulation.GoLibErrorInPackage1 = false
	simulation.GoLibErrorInPackage3 = false
	simulation.CustomErrorInPackage2 = false
}
