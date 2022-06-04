# golang-error-chains
A simple example, demonstrating error-chains in Go.

### Overview

The example demonstrates
- how errors are wrapped
- how errors can be differed at the top (in `main()` function)
- what happens when one package is using another package and all "throw" some errors

### Scenario

The example includes 3 very simple packages:
- package1
- package2
- package3

In our example scenario all 3 packages have some dedicated usage:
- package1 shall read a text file
- package2 shall parse the content of that text file
- package3 shall run some exe file and shall use the parsed result of package2 as arguments for that exe

In the example 3 different errors can occur:
- Some error from using Go standard library, in package1, when reading the file.
- Some own defined custom error, "thrown" in package2, when parsing.
- Some error from using Go standard library, in package3, when running the exe.

The example shows in the `main()` function how to differ the errors at the top. The `main()` function does some special printing, if one of the chained errors is a self-defined custom error (defined in the example code, in the "apperror" package). But if there is some error from the Go standard library, the `main`function does another type of printing.

Everything else is commented in the source code and should be rather self-explanatory. The example is really simple and everyone should get it, just by reading the source code. 😉

#### Have fun.
