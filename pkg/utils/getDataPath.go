package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetDataPath(part string, envvar string) (string, error) {
	hostType, ok := os.LookupEnv(envvar)
	if !ok {
		return "", errors.New(fmt.Sprintf("Environment variabel %s er ikke fundet", envvar))
	}
	if hostType == "kvm" {
		return filepath.Join("/nfs/data", part), nil
	} else if hostType == "none" {
		return filepath.Join("/home/projects/devops/data", part), nil
	} else if hostType == "" {
		return "", errors.New("Envvar er fundet men der er ikke angivet en værdi")
	} else {
		return "", errors.New("Denne host type er ikke understøttet")
	}
}
