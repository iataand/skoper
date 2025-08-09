package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/iataand/skoper/internal/hackerone"
)

func InsertProgram(db *sql.DB, handleData hackerone.Program) (string, error) {
	query := `
        INSERT INTO programs (handle, handleApiUrl)
        VALUES ($1, $2)
        ON CONFLICT (handle) DO UPDATE SET handle=EXCLUDED.handle 
        RETURNING id;
    `

	var id string
	err := db.QueryRow(query, handleData.Handle, handleData.HandleApiUrl).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert program: %w", err)
	}

	return id, nil
}

func InsertScope(db *sql.DB, scope hackerone.Scope, programID string) error {
	createdAt, err := time.Parse(time.RFC3339, scope.Attributes.CreatedAt)
	if err != nil {
		log.Printf("Warning: failed to parse created_at for scope %s: %v", scope.ID, err)
		createdAt = time.Time{}
	}

	updatedAt, err := time.Parse(time.RFC3339, scope.Attributes.UpdatedAt)
	if err != nil {
		log.Printf("Warning: failed to parse updated_at for scope %s: %v", scope.ID, err)
		updatedAt = time.Time{}
	}

	_, err = db.Exec(`
        INSERT INTO scopes (
            id, program_id, asset_type, asset_identifier,
            eligible_for_bounty, eligible_for_submission,
            instruction, max_severity,
            created_at, updated_at,
            confidentiality_requirement, integrity_requirement, availability_requirement
        ) VALUES (
            $1, $2, $3, $4,
            $5, $6,
            $7, $8,
            $9, $10,
            $11, $12, $13
        )
        ON CONFLICT (id) DO NOTHING;
    `,
		scope.ID,
		programID,
		scope.Attributes.AssetType,
		scope.Attributes.AssetIdentifier,
		scope.Attributes.EligibleForBounty,
		scope.Attributes.EligibleForSubmission,
		scope.Attributes.Instruction,
		scope.Attributes.MaxSeverity,
		createdAt,
		updatedAt,
		scope.Attributes.ConfidentialityRequirement,
		scope.Attributes.IntegrityRequirement,
		scope.Attributes.AvailabilityRequirement,
	)
	return err
}
