package main

import (
	"log"
	"os"
)

var (
	dockerRepo        string
	githubRepo        string
	dockerhubEndpoint string
)

func envSetup() error {
	var ok bool

	dockerRepo, ok = os.LookupEnv("DOCKERHUB_REPO")
	if dockerRepo == "" || !ok {
		log.Fatal("DOCKERHUB_REPO env not set")
	}

	githubRepo, ok = os.LookupEnv("GITHUB_REPO")
	if githubRepo == "" || !ok {
		log.Fatal("GITHUB_REPO env not set")
	}

	dockerhubEndpoint, ok = os.LookupEnv("DOCKERHUB_ENDPOINT")
	if dockerhubEndpoint == "" || !ok {
		log.Fatal("DOCKERHUB_ENDPOINT env not set")
	}

	return nil
}
