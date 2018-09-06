package main

import "testing"

func TestParseEmptyArgs(t *testing.T) {
	// Empty args
	args := []string{}
	err := Run(args)
	if err == nil {
		t.Fatal("Don't return error with the empty args.")
	}
}

func TestParseArgsWithUndefinedFlag(t *testing.T) {
	// undefined flag
	args := []string{"-z", "err"}
	err := Run(args)
	if err == nil {
		t.Fatal("Don't return error with the undefined flag.")
	}
}
