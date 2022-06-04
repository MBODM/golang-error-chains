package package2

import (
	"fmt"

	"example.com/error-chains/apperror"
	"example.com/error-chains/package1"
	"example.com/error-chains/simulation"
)

func GetParams() (string, error) {
	// Lets assume we use file contents, we read with help of package1.
	// Then we parse this file contents, but something goes wrong then.
	fileContent, err := package1.ReadFile()
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package2.GetParams() func: %w", err)
	}
	parsedParams, err := parse(fileContent)
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package2.GetParams() func: %w", err)
	}
	return parsedParams, nil
}

func parse(fileContent string) (string, error) {
	// Lets assume we process given file content here and do some parsing.
	// But our parsing fails. Now we have to "throw" our own custom error.
	// Because at the top, in main() function, we wanna differ the errors.
	// ...
	// Do some fancy parsing stuff here.
	// ...
	// Now parsing fails here ("WTF?") and we return our own custom error:
	if simulation.CustomErrorInPackage2 {
		return "", &apperror.CustomError{Err: nil, Msg: "parsing failed cause of WTF situation"}
	}
	// But if all went fine, return str:
	return "/C echo Hello World (some pseudo text result from successful parsing in package2)", nil
}
