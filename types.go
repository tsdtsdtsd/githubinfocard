package githubinfocard

import "github.com/shurcooL/githubv4"

type Graph struct {
	Repository struct {
		DatabaseID  githubv4.Int
		Name        githubv4.String
		Description githubv4.String
		URL         githubv4.URI
		Owner       struct {
			Name githubv4.String `graphql:"login"`
		}

		Issues struct {
			TotalCount githubv4.Int
		} `graphql:"issues(states:OPEN)"`

		Forks struct {
			TotalCount githubv4.Int
		}

		Stargazers struct {
			TotalCount githubv4.Int
		}

		CommitComments struct {
			Nodes []struct {
				PublishedAt githubv4.DateTime
			}
		} `graphql:"commitComments(last:1)"`

		Releases struct {
			Nodes []struct {
				Name        githubv4.String
				PublishedAt githubv4.DateTime
			}
		} `graphql:"releases(last:1)"`

		Languages struct {
			Nodes []struct {
				Name  githubv4.String
				Color githubv4.String
			}
		} `graphql:"languages(first:10)"`
	} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`

	RateLimit struct {
		Cost      githubv4.Int
		Limit     githubv4.Int
		Remaining githubv4.Int
		ResetAt   githubv4.DateTime
	}
}
