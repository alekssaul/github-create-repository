# GitHub Create Repository tool

Creates a github repo via CLI

## Getting Started

- generate an [API Token for Github.com](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)

- `export GITHUB_TOKEN="YOURGITHUBTOKEN"`

- Build the binary: `go build . -o github-create-repository`

- Run the binary: `github-create-repository -name=$NewRepoName`

