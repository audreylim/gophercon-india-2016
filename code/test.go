package main

import "fmt"

// BEGIN OMIT

var fetchUsername = func() string {
	return "real user"
}

var mockUsername = func() string {
	return "fake user"
}

func main() {
	fetchUsername = mockUsername
	fmt.Printf("%v", fetchUsername())
}

// END OMIT
