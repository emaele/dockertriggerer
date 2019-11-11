package getters

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/emaele/dockertriggerer/constants"
	"gitlab.com/emaele/dockertriggerer/types"
)

//GetDockerHubRepoInfo will return repo info from hub.docker.com
func GetDockerHubRepoInfo(repoName string) (types.DockerHubRepoInfo, error) {
	var repoInfo types.DockerHubRepoInfo

	resp, err := http.Get(fmt.Sprintf(constants.DockerHubInfoEndpoint, repoName))
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

func TriggerBuild(endpoint string) error {
	var client http.Client
	var reader io.Reader

	resp, err := client.Post(endpoint, "", reader)
	if err != nil {
		return err
	}

	var postResult types.BuildTriggerResponse

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&postResult)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	if postResult.State == "Success" {
		return nil
	}

	return errors.New("Unable to trigger build")
}
