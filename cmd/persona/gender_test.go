package main

import (
	"testing"

	"github.com/oakroots/persona"
	"github.com/spf13/cobra"
)

func Test_genderValue_All(t *testing.T) {
	tests := []struct {
		name      string
		initial   *persona.Gender
		input     string
		wantErr   bool
		wantValue persona.Gender
		wantStr   string
		wantType  string
	}{
		{
			name:      "default nil value, String should return 'u'",
			initial:   nil,
			input:     "",
			wantErr:   true,
			wantValue: persona.Fantasy,
			wantStr:   "u",
			wantType:  "gender",
		},
		{
			name:      "set male gender",
			initial:   new(persona.Gender),
			input:     "m",
			wantErr:   false,
			wantValue: persona.Male,
			wantStr:   "m",
			wantType:  "gender",
		},
		{
			name:      "set female gender",
			initial:   new(persona.Gender),
			input:     "f",
			wantErr:   false,
			wantValue: persona.Female,
			wantStr:   "f",
			wantType:  "gender",
		},
		{
			name:      "invalid gender",
			initial:   new(persona.Gender),
			input:     "x",
			wantErr:   true,
			wantValue: persona.Fantasy,
			wantStr:   "?",
			wantType:  "gender",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &genderValue{value: tt.initial}

			err := g.Set(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.initial != nil && *tt.initial != tt.wantValue && !tt.wantErr {
				t.Errorf("value after Set() = %v, want %v", *tt.initial, tt.wantValue)
			}

			if got := g.String(); got != tt.wantStr {
				t.Errorf("String() = %v, want %v", got, tt.wantStr)
			}

			if got := g.Type(); got != tt.wantType {
				t.Errorf("Type() = %v, want %v", got, tt.wantType)
			}
		})
	}
}

func TestGeneratorCmd_RunE_DeterministicOnly_NoSeed(t *testing.T) {
	// Zachowanie i przywrócenie globalnych wartości
	origSeed := seed
	origNum := num
	origDet := deterministic
	defer func() {
		seed = origSeed
		num = origNum
		deterministic = origDet
	}()

	// Kluczowy przypadek: seed = 0, deterministic = true
	seed = 0
	num = 1
	deterministic = true

	cmd := &cobra.Command{}
	if err := generatorCmd.RunE(cmd, []string{}); err != nil {
		t.Errorf("RunE() (deterministic=true, seed=0) zwróciło błąd: %v", err)
	}
}
