package main

import (
	"fmt"
	"log"

	"gitlab.com/emaele/dockertriggerer/getters"
)

func main() {
	dockerRepo, err := getters.GetDockerHubRepoInfo("shitpostingio", "tdlib")
	if err != nil {
		log.Fatalf("Unable to perform request: %s", err.Error())
	}

	githubCommit, err := getters.GetLastCommit("tdlib", "td")
	if err != nil {
		log.Fatalf("Unable to perform request: %s", err.Error())
	}

	res := githubCommit.Commit.Author.Date.Sub(dockerRepo.Results[0].LastUpdated)

	if res.Seconds() < 0 {
		fmt.Printf("Your docker image is up to date.\n")
	} else {
		fmt.Printf("Triggering build...\n")
		// TODO: post request
	}
}
