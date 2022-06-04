# golang-error-chains
A simple example, demonstrating error-chains in Go.

### Overview

The example includes 3 very simple packages:
- package1
- package2
- package3

In the example 3 different errors can occur:
- Some error from using Go standard library, in package1.
- Some own defined custom error, "thrown" in package2.
- Some error from using Go standard library, in package3.

The example demonstrates

- how errors are wrapped
- how errors can be differed at the top (in main() function)

when one package use another package and all "throw" some errors.

Everything else is commented in the source code.

The example is really simple and everyone should get it, just by reading the source code.

#### Have fun.
