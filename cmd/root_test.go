package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func runCmd(t *testing.T, args ...string) (string, error) {
	t.Helper()

	orgiStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("could not create pipe: %v", err)
	}
	os.Stdout = w

	rootCmd.SetArgs(args)
	execErr := rootCmd.Execute()

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	os.Stdout = orgiStdout

	rootCmd.SetArgs(nil)

	return buf.String(), execErr
}

// captureOutput redirects stdout for the duration of f and returns what was printed.
func captureOutput(t *testing.T, f func()) string {
	t.Helper()
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = buf.ReadFrom(r)
		close(done)
	}()

	f()

	_ = w.Close()
	os.Stdout = orig
	<-done

	return buf.String()
}

// splitNonEmptyLines splits s into non-empty, trimmed lines.
func splitNonEmptyLines(s string) []string {
	sc := bufio.NewScanner(strings.NewReader(s))
	var lines []string
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func resetGlobals() {
	seed = 0
	deterministic = false
	gender = ""
	num = 1
}

func TestRoot_WithSeed_ForcesDeterministic_PrintsSeed(t *testing.T) {
	resetGlobals()
	// Given a seed, deterministic should be implicitly true (per new rootCmd)
	seed = 123
	num = 2
	gender = "" // agnostic — tylko liczność linii sprawdzamy

	rootCmd.SetArgs([]string{})

	out := captureOutput(t, func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("execute: %v", err)
		}
	})

	lines := splitNonEmptyLines(out)
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (2 names + seed), got %d\noutput:\n%s", len(lines), out)
	}

	// Last line should be the seed "123"
	if got := lines[len(lines)-1]; got != "18874375" {
		t.Fatalf("expected last line to be seed 18874375, got %q\noutput:\n%s", got, out)
	}
}

func TestRoot_NoSeed_NonDeterministic_PrintsOnlyNames(t *testing.T) {
	resetGlobals()
	// No seed and deterministic=false => should NOT print the seed line.
	num = 1
	rootCmd.SetArgs([]string{})

	out := captureOutput(t, func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("execute: %v", err)
		}
	})

	lines := splitNonEmptyLines(out)
	if len(lines) != 1 {
		t.Fatalf("expected exactly 1 line (one full name), got %d\noutput:\n%s", len(lines), out)
	}

	// Be defensive: make sure this one line is not purely numeric (should be a name)
	if _, err := strconv.Atoi(lines[0]); err == nil {
		t.Fatalf("expected a name, got a numeric line: %q", lines[0])
	}
}

func TestRoot_DeterministicWithoutSeed_PrintsSeedLine(t *testing.T) {
	resetGlobals()
	deterministic = true
	num = 1
	rootCmd.SetArgs([]string{})

	out := captureOutput(t, func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("execute: %v", err)
		}
	})

	lines := splitNonEmptyLines(out)
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines (1 name + 1 seed), got %d\noutput:\n%s", len(lines), out)
	}

	// Last line should be some integer seed
	if _, err := strconv.Atoi(lines[1]); err != nil {
		t.Fatalf("expected last line to be numeric seed, got %q", lines[1])
	}
}

func TestPersona_WithSeedAndDeterministic_PrintsSeedAndIsDeterministic(t *testing.T) {
	out1, err := runCmd(t,
		"--num", "2",
		"--gender", "m",
		"--seed", "123",
		"--deterministic",
	)
	if err != nil {
		t.Fatalf("execute (run1): %v", err)
	}

	out2, err := runCmd(t,
		"--num", "2",
		"--gender", "m",
		"--seed", "123",
		"--deterministic",
	)
	if err != nil {
		t.Fatalf("execute (run2): %v", err)
	}

	lines1 := filteredLines(out1)
	lines2 := filteredLines(out2)

	// We expect 3 lines: 2 names + 1 line with the seed.
	if got, want := len(lines1), 3; got != want {
		t.Fatalf("lines (run1): got %d, want %d\nout:\n%s", got, want, out1)
	}
	if got, want := len(lines2), 3; got != want {
		t.Fatalf("lines (run2): got %d, want %d\nout:\n%s", got, want, out2)
	}

	// The last line is the seed (123).
	if got, want := lines1[2], "18874375"; got != want {
		t.Fatalf("seed line (run1): got %q, want %q\nout:\n%s", got, want, out1)
	}
	if got, want := lines2[2], "18874375"; got != want {
		t.Fatalf("seed line (run2): got %q, want %q\nout:\n%s", got, want, out2)
	}

	// Determinism: the same names for the same parameters.
	if lines1[0] != lines2[0] || lines1[1] != lines2[1] {
		t.Fatalf("non-deterministic output for same seed/args:\nrun1: %q\nrun2: %q", lines1[:2], lines2[:2])
	}
}

func filteredLines(s string) []string {
	raw := strings.Split(s, "\n")
	out := make([]string, 0, len(raw))
	for _, ln := range raw {
		ln = strings.TrimSpace(ln)
		if ln != "" {
			out = append(out, ln)
		}
	}
	return out
}
