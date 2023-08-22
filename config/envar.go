package config

import (
	"log"
	"os"
	"strings"
)

var MeshUatGoOneEndpoint string
var RunMode string
var ServerPort string

func InitEnvironmentVariables() {
	RunMode = strings.TrimSpace(os.Getenv("RUN_MODE"))
	if RunMode == "" {
		RunMode = "DEVELOP"
	}

	ServerPort = strings.TrimSpace(os.Getenv("SERVER_PORT"))
	if ServerPort == "" {
		ServerPort = "8081"
	}

	MeshUatGoOneEndpoint = strings.TrimSpace(os.Getenv("MESH_UAT_GO_ONE_ENDPOINT"))
	if MeshUatGoOneEndpoint == "" {
		MeshUatGoOneEndpoint = "http://localhost:8080"
	}

	log.Println("Run Mode: " + RunMode)
	log.Println("Server Port: " + ServerPort)
	log.Println("MeshUatGoOne Endpoint: " + MeshUatGoOneEndpoint)
}
