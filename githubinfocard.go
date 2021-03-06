package githubinfocard

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	// "github.com/davecgh/go-spew/spew"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const (
	githubPrefix = "https://github.com/"
)

// Card defines a repository info card
type Card struct {
	Forks       int
	LastRelease string
	Languages   []Language
	Name        string
	OpenIssues  int
	Owner       string
	URL         string
	Stars       int
}

// Language defines a single programming language
type Language struct {
	Name  string
	Color string
}

// Load tries to load a Card for given repository URL
func Load(url, token string) (*Card, int, error) {

	c := Card{
		URL: url,
	}

	if !strings.HasPrefix(url, githubPrefix) {
		return &c, 0, fmt.Errorf("URL has to start with \"%s\"", githubPrefix)
	}

	// Extract owner and repo name from url
	var err error
	c.Owner, c.Name, err = parseURL(url)
	if err != nil {
		return &c, 0, err
	}

	// API call
	graph, err := fetchGraph(token, c.Owner, c.Name)
	if err != nil {
		return &c, 0, err
	}
	// printJSON(graph)

	// Extract from graph
	c.Forks = int(graph.Repository.Forks.TotalCount)
	c.OpenIssues = int(graph.Repository.Issues.TotalCount)
	c.Stars = int(graph.Repository.Stargazers.TotalCount)

	if len(graph.Repository.Releases.Nodes) >= 1 {
		c.LastRelease = string(graph.Repository.Releases.Nodes[0].Name)
	}

	for _, lang := range graph.Repository.Languages.Nodes {
		c.Languages = append(c.Languages, Language{
			Name:  string(lang.Name),
			Color: string(lang.Color),
		})
	}

	return &c, int(graph.RateLimit.Remaining), nil
}

func parseURL(url string) (string, string, error) {

	tmp := strings.Replace(url, githubPrefix, "", -1)
	tmp = strings.TrimRight(tmp, "/")
	parts := strings.Split(tmp, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("URL invalid, expected \"ownerName/repoName\", got \"%s\"", tmp)
	}

	if parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("URL invalid, expected \"ownerName/repoName\", got \"%s\"", tmp)
	}

	return parts[0], parts[1], nil
}

func fetchGraph(token, owner, repo string) (*Graph, error) {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	gqlVars := map[string]interface{}{
		"repositoryOwner": githubv4.String(owner),
		"repositoryName":  githubv4.String(repo),
	}

	var graph Graph
	err := client.Query(context.Background(), &graph, gqlVars)
	if err != nil {
		return &graph, fmt.Errorf("could not query GitHub API: \"%s\"", err.Error())
	}

	return &graph, nil
}

// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
// only for testing.
func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
