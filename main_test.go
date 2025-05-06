package main

import (
	"testing"
)

func TestTypes(t *testing.T) {
	// Arrange
	cat := Cat{}
	dog := Dog{}
	person := Person{name: "John"}

	expectedOutputs := []string{"Meow", "Woof", "Hello, my name is John"}

	// Act and Assert
	speakers := []Speaker{cat, dog, person}
	for i, speaker := range speakers {
		output := speaker.Speak()
		if output != expectedOutputs[i] {
			t.Errorf("Expected: %s, Got: %s", expectedOutputs[i], output)
		}
	}

	// Additional type checks
	if _, ok := interface{}(cat).(Speaker); !ok {
		t.Errorf("Cat does not implement Speaker")
	}
	if _, ok := interface{}(dog).(Speaker); !ok {
		t.Errorf("Dog does not implement Speaker")
	}
	if _, ok := interface{}(person).(Speaker); !ok {
		t.Errorf("Person does not implement Speaker")
	}
}
