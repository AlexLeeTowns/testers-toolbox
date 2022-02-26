/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
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
		result, err := l.ReadLoremByWordCount(os.DirFS("./pkg/loremipsum"), "lorem.txt", 10)
		if err != nil {
			return err
		}

		e := clipboard.WriteAll(result)

		if e != nil {
			return e
		}
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
}
