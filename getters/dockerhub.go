package getters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/emaele/dockertriggerer/constants"
	"gitlab.com/emaele/dockertriggerer/types"
)

//GetDockerHubRepoInfo will return repo info from hub.docker.com
func GetDockerHubRepoInfo(org, repoName string) (types.DockerHubRepoInfo, error) {
	var repoInfo types.DockerHubRepoInfo

	resp, err := http.Get(fmt.Sprintf(constants.DockerHubInfoEndpoint, org, repoName))
	if err != nil {
		//TODO: Handle error
	}

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&repoInfo)
	if err != nil {
		//TODO: Handle error
	}

	err = resp.Body.Close()
	if err != nil {
		//TODO: Handle error
	}

	return repoInfo, nil
}
