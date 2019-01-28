package githubinfocard_test

import (
	"testing"

	"github.com/tsdtsdtsd/githubinfocard"
)

var repoURL = "https://github.com/tsdtsdtsd/githubinfocard"

func TestLoad(t *testing.T) {

	validCard := githubinfocard.Card{
		URL: repoURL,
	}

	card, err := githubinfocard.Load(repoURL)

	if err != nil {
		t.Error("could not load card:", err)
	}

	if *card != validCard {
		t.Error("invalid card loaded, expected:", validCard, "got:", *card)
	}
}
