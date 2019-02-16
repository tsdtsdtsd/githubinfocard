package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tsdtsdtsd/githubinfocard"
)

var (
	flagURL string
)

func main() {

	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatal("missing mandatory environment variable \"GITHUB_TOKEN\"")
	}

	flag.StringVar(&flagURL, "url", "", "Complete GitHub-URL (e.g. https://github.com/username/reponame)")
	flag.Parse()

	if flagURL == "" {
		log.Fatal("missing mandatory parameter \"--url\"")
	}

	card, remaining, err := githubinfocard.Load(flagURL, os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		log.Fatal("error while loading infocard: " + err.Error())
	}

	fmt.Printf("\n### GitHub Infocard for %s ###\n", card.URL)
	fmt.Printf("Repository: %s\n", card.Name)
	fmt.Printf("Owner: %s\n", card.Owner)

	fmt.Printf("Forks: %d\n", card.Forks)
	fmt.Printf("Stars: %d\n", card.Stars)
	fmt.Printf("Open issues: %d\n", card.OpenIssues)

	fmt.Printf("Last release: %s\n", card.LastRelease)
	fmt.Println("Languages:")
	for _, lang := range card.Languages {
		fmt.Printf(" - %s\n", lang.Name)
	}

	var prio string
	switch {
	case remaining <= 100:
		prio = "! "
	case remaining <= 50:
		prio = "!! "
	case remaining <= 20:
		prio = "!!! "
	}
	fmt.Printf("\n%s(Remaining API calls: %d)\n", prio, remaining)
}
