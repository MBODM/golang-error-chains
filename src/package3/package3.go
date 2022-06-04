package package3

import (
	"fmt"
	"os/exec"
	"strings"

	"example.com/error-chains/package2"
	"example.com/error-chains/simulation"
)

func Execute() (string, error) {
	// Lets assume we run some command line tool with some parameters here.
	// And lets say we get these parameters from our package2 parsing code.
	// Possibility 1: We get an err, containing a Go lib err from package1.
	// Possibility 2: We get our own custom error, from our package2 stuff.
	// Possibility 3: Tool is not installed and we get some Go os/exec err.
	params, err := package2.GetParams()
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package3.Execute() func: %w", err)
	}
	exeFile := "cmd.exe"
	if simulation.GoLibErrorInPackage3 {
		// To create an error here, we just run a non-existing exe file.
		exeFile = "nonexisting.exe"
	}
	output, err := exec.Command(exeFile, params).Output()
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package3.Execute() func: %w", err)
	}
	outputString := string(output)
	outputString = strings.NewReplacer("\"", "", "\r", "", "\n", "").Replace(outputString)
	outputString = strings.TrimSpace(outputString)
	return outputString, nil
}
