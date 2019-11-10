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
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name     string `json:"name"`
		FullSize int    `json:"full_size"`
		Images   []struct {
			Size         int         `json:"size"`
			Digest       string      `json:"digest"`
			Architecture string      `json:"architecture"`
			Os           string      `json:"os"`
			OsVersion    interface{} `json:"os_version"`
			OsFeatures   string      `json:"os_features"`
			Variant      interface{} `json:"variant"`
			Features     string      `json:"features"`
		} `json:"images"`
		ID                  int         `json:"id"`
		Repository          int         `json:"repository"`
		Creator             int         `json:"creator"`
		LastUpdater         int         `json:"last_updater"`
		LastUpdaterUsername string      `json:"last_updater_username"`
		ImageID             interface{} `json:"image_id"`
		V2                  bool        `json:"v2"`
		LastUpdated         time.Time   `json:"last_updated"`
	} `json:"results"`
}
