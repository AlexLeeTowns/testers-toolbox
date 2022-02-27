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

// loremCmd represents the lorem command
var loremCmd = &cobra.Command{
	Use:   "lorem",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var result string

		cmd.Flags().Visit(func(f *pflag.Flag) {
			c, err := strconv.Atoi(f.Value.String())
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			switch f.Name {
			case "char":
				fml, err := l.ReadLoremByCharacterCount(os.DirFS("./pkg/loremipsum"), "lorem.txt", c)
				if err != nil {
					fmt.Print(err)
					os.Exit(1)
				}
				result += fml
			case "paragraph":
				fml, err := l.ReadLoremByParagraph(os.DirFS("./pkg/loremipsum"), "lorem.txt", c)
				if err != nil {
					fmt.Print(err)
					os.Exit(1)
				}
				result += fml
			case "word":
				fml, err := l.ReadLoremByWordCount(os.DirFS("./pkg/loremipsum"), "lorem.txt", c)
				if err != nil {
					fmt.Print(err)
					os.Exit(1)
				}
				result += fml
			}
		})

		clipboard.WriteAll(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loremCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loremCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loremCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loremCmd.Flags().IntP("word", "w", 0, "Retrieve lorem ipsum by word count. If --count is not set, will return 4 words.")
	loremCmd.Flags().IntP("char", "c", 0, "Retrieve lorem ipsum by character count. If --count is not set, will return 5 letters")
	loremCmd.Flags().IntP("paragraph", "p", 0, "Retrieve lorem ipsum by paragraph. If --count is not set, will return 1 paragraph")
}
