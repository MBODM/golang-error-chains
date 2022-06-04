package main

import (
	"errors"
	"fmt"
	"os"

	"example.com/error-chains/app"
	"example.com/error-chains/package3"
	"example.com/error-chains/sim"
)

func main() {
	fmt.Println()
	fmt.Println("Welcome to the error-chains example")
	// Play around with this settings to see how things go:
	sim.SimulateErrorInPackage1 = false
	sim.SimulateErrorInPackage2 = false
	sim.SimulateErrorInPackage3 = false
	// Starting at the top (of the "func call stack") here:
	result, err := package3.Execute()
	if err != nil {
		// We differ the errors here, because we wanna do some
		// special printing, for our self-defined custom error:
		var customError *app.CustomError
		if errors.As(err, &customError) {
			fmt.Println("[Foo] Application Error -->", customError)
		} else {
			fmt.Println("[Bar] Unexpected error(s) occurred -->", err)
		}
		fmt.Println()
		os.Exit(1)
	}
	fmt.Println(result)
	fmt.Println("Have a nice day.")
	fmt.Println()
	os.Exit(0)
}
