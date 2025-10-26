package main

import (
	"fmt"

	"github.com/oakroots/persona"
	"github.com/spf13/cobra"
)

var (
	gender        persona.Gender
	seed          uint32
	num           int
	deterministic bool
)

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var opts []persona.Option
		if seed != 0 {
			opts = append(opts, persona.WithSeed(seed), persona.WithDeterministic())
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

func init() {
	rootCmd.AddCommand(generatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
