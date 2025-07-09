package main

import (
	"encoding/json"
	"log"

	"github.com/iataand/skoper/internal/db"
	"github.com/iataand/skoper/internal/hackerone"
	"github.com/iataand/skoper/internal/utils"
)

func main() {
	inputHandles := utils.ReadTargetsFromStdin()
	user, apiKey := utils.LoadEnvVariables()

	database, err := db.InitDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	for _, handleData := range inputHandles {
		var id string
		id, err = db.InsertProgram(database, handleData)
		if err != nil {
			log.Fatalf("Failed to insert handle into Programs: %v", err)
		}

		scopes, err := hackerone.FetchStructuredScopes(user, apiKey, handleData.HandleApiUrl)
		if err != nil {
			log.Fatal(err)
		}

		var result hackerone.ScopeResponse
		err = json.Unmarshal(scopes, &result)
		if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		for _, scope := range result.Data {
			newScope := hackerone.Scope{
				ID:         scope.ID,
				Attributes: scope.Attributes,
			}

			err := db.InsertScope(database, newScope, id)
			if err != nil {
				log.Printf("Failed to insert scope %s: %v", scope.ID, err)
			}
		}
	}

}
