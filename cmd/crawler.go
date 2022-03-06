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
		url, _ := cmd.Flags().GetString("url")
		d, _ := cmd.Flags().GetInt("depth")
		crawler.Crawl(os.Stdout, url, d, &f)

		if e != nil {
			return e
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)
	crawlerCmd.Flags().IntP("depth", "d", 1, "Define how many recursions the page should check")
	crawlerCmd.Flags().StringP("url", "u", "http://localhost:8080", "Define URL startingpoint to crawl")
}
