# github-infocard

> Experimental Go library which fetches summaries of GitHub repositories

[![Build Status][travis-image]][travis-url]

## Idea

I wanted to create an element on my blog, which displays a summary for a given GitHub repository. It could show visitors information like:

- the repository owners name
- amount of stars, forks and open issues
- date of last commit
- maybe even the language details (the narrow colored bar above the branch button)

Searching for ways to fetch this kind of data I found the GitHub API and specificaly it's [GraphQL API](https://developer.github.com/v4/guides/intro-to-graphql/).
As I never messed with GraphQL, I thought that this is the perfect opportunity to learn some new things.

My current idea is to develop a library that helps fetching the data. <br>
As an example and testing application, I want to add a CLI client, which just shows the desired data in an extraordinarily unspectacular manner.

Oh, and this is a plan-as-you-develop kind of project ...

## GitHub API

You need a *Personal access token* for the API, which you can [generate here](https://github.com/settings/tokens) (you need to be signed in first).

## Usage

```
go get github.com/tsdtsdtsd/githubinfocard 
```

There is an example CLI client in `example/cli/`. It needs two things from you:
- Environment variable named `GITHUB_TOKEN`, containing your personal access token
- Parameter `--url` to indicate a repository

The following commands assume that you are in `$GO_PATH/github.com/tsdtsdtsd/githubinfocard`:

### Windows

```
cmd /V /C "set GITHUB_TOKEN=##TOKEN## && go run example\cli\cli.go --url https://github.com/torvalds/linux"
```

### Linux

```
GITHUB_TOKEN=##TOKEN## bash -c 'go run example/cli/cli.go --url https://github.com/torvalds/linux
```

<!-- Markdown link & img dfn's -->
[travis-image]: https://travis-ci.org/tsdtsdtsd/githubinfocard.svg?branch=master
[travis-url]: https://travis-ci.org/tsdtsdtsd/githubinfocard