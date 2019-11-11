package types

import "time"

//BuildTriggerResponse contains selected fields of the POST request response
type BuildTriggerResponse struct {
	State string `json:"state"`
	UUID  string `json:"uuid"`
}

// GithubCommit contains selected fields of https://api.github.com/repos/org/reponame/commits/master api call
type GithubCommit struct {
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		}
	} `json:"commit"`
}

// DockerHubRepoInfo contains selected fields of the docker info api call
type DockerHubRepoInfo struct {
	Name        string    `json:"name"`
	ID          int       `json:"id"`
	Repository  int       `json:"repository"`
	LastUpdated time.Time `json:"last_updated"`
}
