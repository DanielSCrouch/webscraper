package main

import (
	"fmt"
	"log"
	"os"
	"scraper/pkg/cmdcli"
	"scraper/pkg/scraper"
)

func main() {
	logger := log.Default()
	cli := cmdcli.NewCLI("vidsy", *logger)

	cli.RegisterCmd("scrape", scraperCmd)

	cli.Run()
}

func scraperCmd(args []string) {
	if len(args) != 1 {
		fmt.Printf("invalid argument length %d\n", len(args))
		os.Exit(1)
	}

	uri := args[0]
	htmlScraper := scraper.NewScraper(scraper.WebPageFormat("html"))
	htmlScraper.Scrape(uri)

	attributes, err := htmlScraper.ParseAttributes("href", "a")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	for _, attr := range attributes {
		fmt.Println(attr)
	}

}
