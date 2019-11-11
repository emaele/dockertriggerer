package getters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/emaele/dockertriggerer/constants"
	"gitlab.com/emaele/dockertriggerer/types"
)

//GetLastCommit will return the last master commit
func GetLastCommit(repoName string) (types.GithubCommit, error) {
	var lastCommit types.GithubCommit

	resp, err := http.Get(fmt.Sprintf(constants.GithubCommitEndpoint, repoName))
	if err != nil {
		return lastCommit, err
	}

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&lastCommit)
	if err != nil {
		return lastCommit, err
	}

	err = resp.Body.Close()
	if err != nil {
		return lastCommit, err
	}

	return lastCommit, nil
}
