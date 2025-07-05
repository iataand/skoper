package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/iataand/skoper/internal/db"
	"github.com/iataand/skoper/internal/hackerone"
	"github.com/iataand/skoper/internal/utils"
)

func main() {
	handles := utils.ReadTargetsFromStdin()
	user, apiKey := utils.LoadEnvVariables()

	database, err := db.InitDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	for _, handle := range handles {
		var id string
		id, err = db.InsertProgram(database, handle)
		if err != nil {
			log.Fatalf("Failed to insert handle into Programs: %v", err)
		}

		fmt.Println(id)

		scopes, err := hackerone.FetchStructuredScopes(user, apiKey, handle)
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
