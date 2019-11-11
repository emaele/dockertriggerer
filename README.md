# DockerHub Build Triggerer

A simple tool that triggers a DockerHub build if your build in latest is older than the last commit on master of a GitHub repository.

## How to run

To run the tool you have to set up some config env vars as following:

| **Var** 	            | **Description**                 	                                    |
|-----------------------|-----------------------------------------------------------------------|
| `DOCKERHUB_REPO`      | Your DockerHub repo in the format `user/reponame` or `org/reponame`   |
| `GITHUB_REPO`      	| Your GitHub repo in the format `user/reponame` or `org/reponame`      |
| `DOCKERHUB_ENDPOINT`  | Is the endpoint given to you by DockerHub when you set a build trigger|

After done that, build the tool and execute it when you want with a cronjob or a systemd timer.
