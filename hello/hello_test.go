package main

import "testing"

// Writing tests
// Writing a test is just like writing a function, with a few rules
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

func TestHello(t *testing.T) {
	// refactor to minimize repeition:
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper needed to tell the test suit that this method is a helper
		//  this allows for function specific error messages and not this helper function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Greg")
		want := "Hello, Greg"
		assertCorrectMessage(t, got, want)
	})
	// Subtests
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
