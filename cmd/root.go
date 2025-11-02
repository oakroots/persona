package main

import (
	"fmt"

	"github.com/oakroots/persona"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "persona",
	Short: "Generate fake name and surname",
	RunE: func(cmd *cobra.Command, args []string) error {
		var opts []persona.Option
		if seed != 0 {
			opts = append(opts, persona.WithSeed(seed), persona.WithDeterministic())
			deterministic = true
		} else if deterministic {
			opts = append(opts, persona.WithDeterministic())
		}

		opts = append(opts, persona.WithGender(gender))

		pg := persona.New(opts...)

		for i := 0; i < num; i++ {
			fmt.Printf("%s\n", pg.GetFullName())
		}

		if deterministic {
			fmt.Println(pg.Seed())
		}

		return nil
	},
}
