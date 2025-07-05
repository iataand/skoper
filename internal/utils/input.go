package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Target struct {
	Handle       string
	HandleApiURL string
}

func ReadTargetsFromStdin() []Target {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin")
	}

	lines := strings.Split(string(data), "\n")
	var resultList []Target

	for _, line := range lines {
		handle := strings.TrimSpace(line)
		if handle != "" {
			url := fmt.Sprintf("https://api.hackerone.com/v1/hackers/programs/%s/structured_scopes", handle)
			resultList = append(resultList, Target{
				Handle:       handle,
				HandleApiURL: url,
			})
		}
	}

	return resultList
}
