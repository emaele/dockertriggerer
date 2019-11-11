package constants

//DockerHubInfoEndpoint is the endpoint for retrieve info about a docker repository
const DockerHubInfoEndpoint = "https://hub.docker.com/v2/repositories/%s/tags/latest"

//GithubCommitEndpoint is the api endpoint for the last commit of a github repo
const GithubCommitEndpoint = "https://api.github.com/repos/%s/commits/master"
