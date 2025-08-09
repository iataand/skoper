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
		scopes, err := hackerone.FetchStructuredScopes(user, apiKey, handleData.HandleApiUrl)
		if err != nil {
			log.Printf("Failed to fetch scopes for handle %s: %v", handleData.Handle, err)
			continue
		}

		var result hackerone.ScopeResponse
		err = json.Unmarshal(scopes, &result)
		if err != nil {
			log.Printf("Failed to parse JSON for handle %s: %v", handleData.Handle, err)
			continue
		}

		var id string
		id, err = db.InsertProgram(database, handleData)
		if err != nil {
			log.Printf("Failed to insert handle %s into Programs: %v", handleData.Handle, err)
			continue
		}

		log.Printf("Successfully inserted program %s with %d scopes", handleData.Handle, len(result.Data))

		for _, scope := range result.Data {
			newScope := hackerone.Scope{
				ID:         scope.ID,
				Attributes: scope.Attributes,
			}

			err := db.InsertScope(database, newScope, id)
			if err != nil {
				log.Printf("Failed to insert scope %s for handle %s: %v", scope.ID, handleData.Handle, err)
			}
		}
	}

}
