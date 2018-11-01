package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// RepositoryRequest - Holds New Repository request
type RepositoryRequest struct {
	Name              string
	GithubAccessToken string
	AutoInit          bool
}

func main() {
	reponame := flag.String("name", "name", "Name of Repo to create")
	githubtoken := os.Getenv("GITHUB_TOKEN")
	flag.Parse()

	if *reponame == "" {
		fmt.Printf("Please enter name ie:\"-name=newrepo\"\n")
		os.Exit(1)
	}

	if githubtoken == "" {
		fmt.Printf("Please export the environmental variable GITHUB_TOKEN\n")
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubtoken},
	)
	tc := oauth2.NewClient(ctx, ts)

	GithubClient := github.NewClient(tc)
	autoinit := true
	newrepo := github.Repository{
		Name:     reponame,
		AutoInit: &autoinit,
	}

	log.Printf("Creating a repository %v\n", newrepo)
	_, response, err := GithubClient.Repositories.Create(ctx, "", &newrepo)
	if err != nil {
		log.Panicf("Error Occured %v", err)
	}
	fmt.Printf("Result: %v\n", response.StatusCode)

}
