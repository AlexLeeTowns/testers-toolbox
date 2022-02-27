/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
	"github.com/atotto/clipboard"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func readLorem(name string, count int) string {
	switch name {
	case "char":
		fml, err := l.ReadLoremByCharacterCount(os.DirFS("./pkg/loremipsum"), "lorem.txt", count)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		return fml
	case "paragraph":
		fml, err := l.ReadLoremByParagraph(os.DirFS("./pkg/loremipsum"), "lorem.txt", count)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		return fml
	case "word":
		fml, err := l.ReadLoremByWordCount(os.DirFS("./pkg/loremipsum"), "lorem.txt", count)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		return fml
	}

	return ""
}

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
			result += readLorem(f.Name, c)
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
	loremCmd.Flags().IntP("char", "c", 0, "Retrieve lorem ipsum by character count. If --count is not set, will return 5 letters")
	loremCmd.Flags().IntP("paragraph", "p", 0, "Retrieve lorem ipsum by paragraph. If --count is not set, will return 1 paragraph")
}
