/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
	"github.com/atotto/clipboard"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "lorem",
	Short: "Copy lorem ipsum text to clipboard",
	Long:  `Copy lorem ipsum text to clipboard by word count, character count, or paragraph count.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var result string
		var e error

		cmd.Flags().Visit(func(f *pflag.Flag) {
			c, err := strconv.Atoi(f.Value.String())
			if err != nil {
				e = err
			}
			loremText := l.Read(f.Name, c)

			result += loremText
		})

		if e != nil {
			return e
		}

		clipboard.WriteAll(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loremCmd)
	loremCmd.Flags().IntP("char", "c", 0, "Copy lorem ipsum text to clipboard by character count.")
	loremCmd.Flags().IntP("paragraph", "p", 0, "Copy lorem ipsum text to clipboard by paragraph count.")
	loremCmd.Flags().IntP("word", "w", 0, "Copy lorem ipsum text to clipboard by word count")
}
