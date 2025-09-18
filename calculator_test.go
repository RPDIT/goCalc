package main

import "testing"

func TestFloatConversion(t *testing.T) {
	expect := float64(12345)
	got := ParseFloatsfromStrings("12345")
	if got != expect {
		t.Error("Failed to convert to float from string.")
	}
}
