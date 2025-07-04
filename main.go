package main

import (
	"encoding/json"
	"fmt"
	"github.com/lpernett/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	urls := readTargetsFromStdin()
	user, apiKey := loadEnvVariables()

	for _, url := range urls {
		scopes, err := fetchStructuredScopes(user, apiKey, url)
		if err != nil {
			log.Fatal(err)
		}

		var result ScopeResponse
		err = json.Unmarshal(scopes, &result)
		if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		for _, scope := range result.Data {
			fmt.Printf(
				"ID: %s, Url: %s, Bounty: %t\n",
				scope.ID,
				scope.Attributes.AssetIdentifier,
				scope.Attributes.EligibleForBounty,
			)
		}
	}

}

func fetchStructuredScopes(user, apiKey, baseURL string) ([]byte, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	query := parsedURL.Query()
	query.Set("page[size]", "100")
	parsedURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(user, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}

func readTargetsFromStdin() []string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin:")
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

func loadEnvVariables() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("HACKERONE_USERNAME")
	apiKey := os.Getenv("HACKERONE_API_KEY")

	if user == "" || apiKey == "" {
		log.Fatal("Missing HACKERONE_USERNAME or HACKERONE_API_KEY in environment")
	}

	return user, apiKey
}
