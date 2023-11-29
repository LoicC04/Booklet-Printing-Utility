package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestBrochureOrderedPage20(t *testing.T) {
	// GIVEN
	nbPage := 20
	expected := "20,1,2,19,18,3,4,17,16,5,6,15,14,7,8,13,12,9,10,11"

	// WHEN
	result := BookletOrderedPage(nbPage)

	// THEN
	if result != expected {
		t.Fatalf(`Expected '%v' but got '%v'`, expected, result)
	}
}

func TestBrochureOrderedPage4(t *testing.T) {
	// GIVEN
	nbPage := 4
	expected := "4,1,2,3"

	// WHEN
	result := BookletOrderedPage(nbPage)

	// THEN
	if result != expected {
		t.Fatalf(`Expected '%v' but got '%v'`, expected, result)
	}
}

func TestBrochureOrderedPage8(t *testing.T) {
	// GIVEN
	nbPage := 8
	expected := "8,1,2,7,6,3,4,5"

	// WHEN
	result := BookletOrderedPage(nbPage)

	// THEN
	if result != expected {
		t.Fatalf(`Expected '%v' but got '%v'`, expected, result)
	}
}
