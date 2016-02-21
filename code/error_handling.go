package main

import "fmt"

// BEGIN OMIT

func mailErrorLogger() func(error) {
	return func(err error) {
		fmt.Printf("Failed to send mail: %v", err)
	}
}

func doSomething(errHandler func(error)) {
	if err := sendMail(); err != nil {
		errHandler(err)
	}
}

func sendMail() error {
	return fmt.Errorf("an error occurred")
}

func main() {
	errHandler := mailErrorLogger()

	doSomething(errHandler)
}

// END OMIT
