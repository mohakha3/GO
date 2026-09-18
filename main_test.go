package main

import "testing"

func TestRunAppStandard(t *testing.T) {
	got := RunApp("Cisco Vet", false)
	want := "Hello, Cisco Vet."
	if got != want { t.Errorf("got %q; want %q", got, want) }
}

func TestRunAppVerbose(t *testing.T) {
	got := RunApp("Mag7 Architect", true)
	want := "[DEBUG ENGINE ACTIVE] Welcome Senior Leader, Mag7 Architect."
	if got != want { t.Errorf("got %q; want %q", got, want) }
}
