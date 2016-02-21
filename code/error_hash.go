package main

import "fmt"

// BEGIN OMIT

func main() {
	errHandler := mailErrorLogger()

	doSomething(errHandler)
}

// END OMIT
