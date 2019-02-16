package githubinfocard_test

import (
	"os"
	"testing"

	"github.com/tsdtsdtsd/githubinfocard"
)

var repoURL = "https://github.com/tsdtsdtsd/githubinfocard"

func TestLoad(t *testing.T) {

	// @todo this needs a fix now that we have githubinfocard.Card.Languages

	// validCard := githubinfocard.Card{
	// 	URL: repoURL,
	// }

	_, _, err := githubinfocard.Load(repoURL, os.Getenv("GITHUB_TOKEN"))

	if err != nil {
		t.Error("could not load card:", err)
	}

	// if *card != validCard {
	// 	t.Error("invalid card loaded, expected:", validCard, "got:", *card)
	// }
}
