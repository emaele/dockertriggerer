package main

import (
	"fmt"
	"log"

	"gitlab.com/emaele/dockertriggerer/getters"
)

func main() {
	dockerRepo, err := getters.GetDockerHubRepoInfo(dockerRepo)
	if err != nil {
		log.Fatalf("Unable to perform request: %s", err.Error())
	}

	githubCommit, err := getters.GetLastCommit(githubRepo)
	if err != nil {
		log.Fatalf("Unable to perform request: %s", err.Error())
	}

	res := githubCommit.Commit.Author.Date.Sub(dockerRepo.LastUpdated)

	if res.Seconds() < 0 {
		fmt.Printf("Your docker image is up to date.\n")
	} else {
		fmt.Printf("Triggering build...\n")
		err := getters.TriggerBuild(dockerhubEndpoint)
		if err != nil {
			log.Fatalf("Unable to trigger build: %s", err.Error())
		}
	}
}
