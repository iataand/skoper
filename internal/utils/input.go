package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/iataand/skoper/internal/hackerone"
)

func ReadTargetsFromStdin() []hackerone.Program {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin")
	}

	lines := strings.Split(string(data), "\n")
	var resultList []hackerone.Program

	for _, line := range lines {
		handle := strings.TrimSpace(line)
		if handle != "" {
			url := fmt.Sprintf("https://api.hackerone.com/v1/hackers/programs/%s/structured_scopes", handle)
			resultList = append(resultList, hackerone.Program{
				Handle:       handle,
				HandleApiUrl: url,
			})
		}
	}

	return resultList
}
