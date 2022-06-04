package package2

import (
	"fmt"

	"example.com/error-chains/app"
	"example.com/error-chains/package1"
	"example.com/error-chains/sim"
)

func GetParams() (string, error) {
	// Lets assume we use file contents, we read with help of package1.
	// Then we parse this file contents, but something goes wrong then.
	fileContent, err := package1.ReadFile()
	if err != nil {
		// We wrap the error we got, into some typical non-custom-error:
		return "", fmt.Errorf("error in package2.GetParams() func: %w", err)
		//
		// Important (for later use):
		//
		// Now lets have a look at what happens when we wrap this error
		// again. But this time into our custom error. So we can differ
		// this situation in main() too, and show some specific message.
		// To test this: Comment out above line and uncomment next line.
		//
		// return "", &app.CustomError{Err: err, Msg: "could not read params from text file"}
	}
	parsedParams, err := parse(fileContent)
	if err != nil {
		return "", err
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
	if sim.SimulateErrorInPackage2 {
		return "", &app.CustomError{Err: nil, Msg: "parsing failed cause of WTF situation"}
	}
	// But if all went fine, return that parsed params (ofc we fake this):
	parsedParams := "/C echo Hello World (some pseudo result from successful parsing in package2)"
	return parsedParams, nil
}
