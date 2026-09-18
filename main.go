package main

import (
	"flag"
	"fmt"
)

// RunApp processes variables using pointers extracted from the flag package
func RunApp(name string, verbose bool) string {
	if verbose {
		return fmt.Sprintf("[DEBUG ENGINE ACTIVE] Welcome Senior Leader, %s.", name)
	}
	return fmt.Sprintf("Hello, %s.", name)
}

func main() {
	// flag.String returns a *string (pointer), NOT a plain string value
	namePtr := flag.String("name", "Guest", "The name of the operator")
	verbosePtr := flag.Bool("verbose", false, "Enable heavy debug logging")

	// This is the critical engine step—it parses the terminal arguments array
	flag.Parse()

	// We use the '*' operator to "dereference" the pointers and read the true value
	result := RunApp(*namePtr, *verbosePtr)
	fmt.Println(result)
}
