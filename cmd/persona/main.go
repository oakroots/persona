package main

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.Flags().VarP(&genderValue{&gender}, "gender", "g", "Gender: f=female, m=male, u=fantasy")
	rootCmd.Flags().IntVarP(&num, "num", "n", 1, "Number of personas")
	rootCmd.Flags().Uint32VarP(&seed, "seed", "s", 0, "Pseudo-random seeded sequence (1 to 2^25-1)")
	rootCmd.Flags().BoolVarP(&deterministic, "deterministic", "d", false, "Deterministic seed")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
