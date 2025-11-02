package main

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

// Test scenariusza gdy rootCmd.Execute() nie zwraca błędu.
func TestMain_NoError(t *testing.T) {
	// Blok warunkowy wykonywany **tylko w procesie potomnym**:
	if os.Getenv("TEST_MAIN_SUCCESS") == "1" {
		// Uruchomienie main() w procesie potomnym.
		main()
		// Zakończenie procesu potomnego po wykonaniu main.
		os.Exit(0)
	}

	// Konfiguracja i uruchomienie procesu potomnego w procesie testowym (macierzystym):
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_NoError")
	cmd.Env = append(os.Environ(), "TEST_MAIN_SUCCESS=1")
	output, err := cmd.CombinedOutput()

	// Sprawdzenie kodu wyjścia procesu potomnego.
	exitCode := getExitCode(err)
	if exitCode != 0 {
		t.Fatalf("Proces z main powinien zakończyć się kodem 0, a zwrócił %d (błąd: %v)", exitCode, err)
	}

	// Sprawdzenie, że na wyjściu NIE pojawił się komunikat "Error: ...".
	if bytes.Contains(output, []byte("Error:")) {
		t.Errorf("Nieoczekiwany komunikat błędu na wyjściu: %q", output)
	}
}

// Test scenariusza gdy rootCmd.Execute() zwraca błąd.
func TestMain_Error(t *testing.T) {
	if os.Getenv("TEST_MAIN_ERROR") == "1" {
		// Symulacja błędu: przekazujemy nieznaną flagę aby wymusić błąd Execute().
		os.Args = []string{os.Args[0], "--badflag"}
		main()
		os.Exit(0) // Jeśli main nie wywoła os.Exit, zakończ proces potomny samodzielnie.
	}

	// Uruchomienie procesu potomnego z ustawioną flagą testową.
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_Error")
	cmd.Env = append(os.Environ(), "TEST_MAIN_ERROR=1")
	output, err := cmd.CombinedOutput()

	// Kod wyjścia powinien wskazywać na błąd (exit status 1).
	exitCode := getExitCode(err)
	if exitCode == 0 {
		t.Fatalf("Oczekiwano kodu wyjścia 1 przy błędzie, a otrzymano %d", exitCode)
	}

	// Sprawdzenie, czy komunikat na stdout zawiera prefix "Error:" i treść błędu z Execute().
	if !bytes.Contains(output, []byte("Error:")) {
		t.Errorf("Brak oczekiwanego komunikatu błędu. Wyjście: %q", output)
	}
	// (Opcjonalnie) Można doprecyzować treść komunikatu, np. "unknown flag" dla --badflag.
	if !bytes.Contains(output, []byte("unknown flag")) {
		t.Errorf("Komunikat błędu nie zawiera treści błędu Execute(): %q", output)
	}
}

// getExitCode ekstraktuje kod wyjścia z błędu procesu (jeśli proces zakończył się niezerowym kodem).
func getExitCode(err error) int {
	if err == nil {
		return 0 // brak błędu oznacza kod wyjścia 0
	}
	// Sprawdzenie typu błędu i ekstrakcja statusu.
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		}
	}
	// Jeśli nie udało się ustalić kodu (lub inny błąd), zwracamy 0 jako domyślny.
	return 0
}
