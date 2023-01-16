package controller

import (
	"go-webscrapper/application"
	"strings"

	"github.com/spf13/cobra"
)

var defaultInput = "https://catracalivre.com.br/feed/,https://www.infomoney.com.br/feed/,https://forbes.com.br/feed/,https://www.cnnbrasil.com.br/feed/,https://www.moneytimes.com.br/feed/"

func NewFeedController(usecase application.CreateFeeds) *cobra.Command {
	feedController := &cobra.Command{
		Use:   "feed",
		Short: "use this controller scrapp pages from news websites(feed pages)",
		Long: `
pass the URL from the news websites and save all of them in a sql database. 
You can pass a flag called input(or the shorthand -i) and passes all the websites separada by comma. 
Example: --input=https://catracalivre.com.br/feed/,https://www.infomoney.com.br/feed/
If you run the program without the flag, it will execute with the default input websites.
`,
		RunE: controller(usecase),
	}

	feedController.Flags().StringP("input", "i", defaultInput, "pass the websites URL separate by comma")
	return feedController
}

func controller(usecase application.CreateFeeds) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		usecase.Execute(strings.Split(input, ","))
		return nil
	}
}
