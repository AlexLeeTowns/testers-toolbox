package cmd

import (
	"os"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/crawler"
	"github.com/spf13/cobra"
)

var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawl webpage looking for 404s",
	RunE: func(cmd *cobra.Command, args []string) error {
		var e error
		f := make(crawler.Fetchy, 0)
		// TODO: Make it actually respond to input
		crawler.Crawl(os.Stdout, "https://ultimateqa.com/automation", 2, &f)

		if e != nil {
			return e
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)
}
