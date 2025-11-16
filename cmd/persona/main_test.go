package main

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestMain_NoError(t *testing.T) {
	if os.Getenv("TEST_MAIN_SUCCESS") == "1" {
		main()
		os.Exit(0)
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_NoError")
	cmd.Env = append(os.Environ(), "TEST_MAIN_SUCCESS=1")
	output, err := cmd.CombinedOutput()

	exitCode := getExitCode(err)
	if exitCode != 0 {
		t.Fatalf("Proces z main powinien zakończyć się kodem 0, a zwrócił %d (błąd: %v)", exitCode, err)
	}

	if bytes.Contains(output, []byte("Error:")) {
		t.Errorf("Nieoczekiwany komunikat błędu na wyjściu: %q", output)
	}
}

func TestMain_Error(t *testing.T) {
	if os.Getenv("TEST_MAIN_ERROR") == "1" {
		os.Args = []string{os.Args[0], "--badflag"}
		main()
		os.Exit(0)
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_Error")
	cmd.Env = append(os.Environ(), "TEST_MAIN_ERROR=1")
	output, err := cmd.CombinedOutput()

	exitCode := getExitCode(err)
	if exitCode == 0 {
		t.Fatalf("Oczekiwano kodu wyjścia 1 przy błędzie, a otrzymano %d", exitCode)
	}

	if !bytes.Contains(output, []byte("Error:")) {
		t.Errorf("Brak oczekiwanego komunikatu błędu. Wyjście: %q", output)
	}
	if !bytes.Contains(output, []byte("unknown flag")) {
		t.Errorf("Komunikat błędu nie zawiera treści błędu Execute(): %q", output)
	}
}

func getExitCode(err error) int {
	if err == nil {
		return 0 // brak błędu oznacza kod wyjścia 0
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		}
	}

	return 0
}
