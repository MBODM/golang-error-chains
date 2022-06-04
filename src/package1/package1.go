package package1

import (
	"fmt"
	"os"

	"example.com/error-chains/simulation"
)

func ReadFile() (string, error) {
	// Lets assume we read some file here, to get input for package2.
	// But the file not exists, so we get some os/exec error from Go.
	fileName := getFileName()
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package1.ReadFile() func: %w", err)
	}
	fileContent := string(bytes)
	return fileContent, nil
}

func getFileName() string {
	if simulation.GoLibErrorInPackage1 {
		// Depending on simulation settings, we force the error by reading a non-existing file.
		return `C:\Windows\non-existing-file.txt`
	} else {
		// If we do not want to create some error here, we just read a file, every Windows has.
		return `C:\Windows\system.ini`
	}
}
