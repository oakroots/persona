package main

import (
	"testing"

	"github.com/oakroots/persona"
	"github.com/spf13/cobra"
)

func TestGeneratorCmd_RunE_Simple(t *testing.T) {
	origSeed := seed
	origNum := num
	origDet := deterministic
	origGender := gender
	defer func() {
		seed = origSeed
		num = origNum
		deterministic = origDet
		gender = origGender
	}()

	seed = 0
	num = 1
	deterministic = false
	gender = persona.Female

	cmd := &cobra.Command{}
	if err := generatorCmd.RunE(cmd, []string{}); err != nil {
		t.Errorf("RunE() zwróciło błąd: %v", err)
	}
}

func TestGeneratorCmd_RunE_DeterministicOnly(t *testing.T) {
	origSeed := seed
	origNum := num
	origDet := deterministic
	defer func() {
		seed = origSeed
		num = origNum
		deterministic = origDet
	}()

	seed = 0
	num = 1
	deterministic = true

	cmd := &cobra.Command{}
	if err := generatorCmd.RunE(cmd, []string{}); err != nil {
		t.Errorf("RunE() dla deterministic zwróciło błąd: %v", err)
	}
}

func TestGeneratorCmd_RunE_WithSeed(t *testing.T) {
	origSeed := seed
	origNum := num
	origDet := deterministic
	defer func() {
		seed = origSeed
		num = origNum
		deterministic = origDet
	}()

	seed = 12345
	num = 1
	deterministic = false

	cmd := &cobra.Command{}
	if err := generatorCmd.RunE(cmd, []string{}); err != nil {
		t.Errorf("RunE() z seed > 0 zwróciło błąd: %v", err)
	}
}
