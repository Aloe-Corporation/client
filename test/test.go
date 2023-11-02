package test

import "os"

var (
	// WorkDir path for the project.
	WorkDir = os.Getenv("GOPATH") + "/src/github.com/Aloe-Corporation/client"

	// DockerCompose .
	DockerCompose = WorkDir + "/test/deployment/docker-compose.yaml"
)
