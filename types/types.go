package types

import "time"

type BuildTriggerResponse struct {
	Autotests     string        `json:"autotests"`
	BuildInFarm   bool          `json:"build_in_farm"`
	BuildSettings []string      `json:"build_settings"`
	Channel       string        `json:"channel"`
	Deploykey     string        `json:"deploykey"`
	Envvars       []interface{} `json:"envvars"`
	Image         string        `json:"image"`
	Owner         string        `json:"owner"`
	Provider      string        `json:"provider"`
	RepoLinks     bool          `json:"repo_links"`
	Repository    string        `json:"repository"`
	ResourceURI   string        `json:"resource_uri"`
	State         string        `json:"state"`
	UUID          string        `json:"uuid"`
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

type DockerHubRepoInfo struct {
	Name        string    `json:"name"`
	ID          int       `json:"id"`
	Repository  int       `json:"repository"`
	LastUpdated time.Time `json:"last_updated"`
}
