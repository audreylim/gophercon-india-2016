package main

import "fmt"

// BEGIN OMIT

type FakeUser struct {
	Username string
}

func (u *FakeUser) FetchUsername() string {
	return "fake user"
}

type ExternalService interface {
	FetchUsername() string
}

func main() {
	fakeExternalService := &FakeUser{}
	fmt.Println(fakeExternalService.FetchUsername())
}

// END OMIT
