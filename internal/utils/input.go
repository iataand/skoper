package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadTargetsFromStdin() []string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin")
	}

	lines := strings.Split(string(data), "\n")
	var resultList []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			url := fmt.Sprintf("https://api.hackerone.com/v1/hackers/programs/%s/structured_scopes", line)
			resultList = append(resultList, url)
		}
	}

	return resultList
}
