package githubinfocard

// Card defines a repository info card
type Card struct {
	Forks      int
	Name       string
	OpenIssues int
	Owner      string
	URL        string
	Stars      int
}

// Load tries to load a Card for given repository URL
func Load(url string) (*Card, error) {
	return &Card{
		URL: url,
	}, nil
}
