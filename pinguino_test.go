package main

import "testing"

func TestLoadconfig(t *testing.T) {
	var res bool
	res = loadconfig()
	if res != true {
		t.Error("Expected true, got ", res)
	}
}

func TestRunCommand(t *testing.T) {
	var res bool
	command := []string{"touch", "file.txt"}
	res = runCommand(command)
	if res != true {
		t.Error("Expected true, got ", res)
	}
	command = []string{""}
	res = runCommand(command)
	if res != true {
		t.Error("Expected true, got ", res)
	}
}
