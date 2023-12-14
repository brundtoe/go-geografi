package utils

import (
	"fmt"
	"os"
)

func OpenDataFile(partFilename string) (*os.File, error) {
	filename, err := GetDataPath(partFilename, "PLATFORM")
	if err != nil {
		return nil, fmt.Errorf("Det er ikke muligt at finde path filen %s: %s", partFilename, err)
	}
	fp, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Det er ikke muligt at åbne filen %s: %s", partFilename, err)
	}

	return fp, nil
}
