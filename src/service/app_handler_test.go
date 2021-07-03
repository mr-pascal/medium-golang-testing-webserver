package main

import (
	"testing"
)

func TestAppHandlerStruct(t *testing.T) {

	t.Run(`Sum`, func(t *testing.T) {
		handler := &AppHandlerStruct{}

		got := handler.Sum(2, 3)
		want := Result{
			Value: 5,
		}

		if got != want {
			t.Errorf("Sum was incorrect, got: %v, want: %v.", got, want)

		}
	})

	t.Run(`Multiply`, func(t *testing.T) {
		handler := &AppHandlerStruct{}

		got := handler.Multiply(2, 3)
		want := Result{
			Value: 6,
		}

		if got != want {
			t.Errorf("Multiply was incorrect, got: %v, want: %v.", got, want)

		}
	})

}
