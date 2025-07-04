package main

type ScopeResponse struct {
	Data  []Data `json:"data"`
	Links Links  `json:"links"`
}

type Attributes struct {
	AssetType                  string `json:"asset_type"`
	AssetIdentifier            string `json:"asset_identifier"`
	EligibleForBounty          bool   `json:"eligible_for_bounty"`
	EligibleForSubmission      bool   `json:"eligible_for_submission"`
	Instruction                string `json:"instruction"`
	MaxSeverity                string `json:"max_severity"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
	ConfidentialityRequirement string `json:"confidentiality_requirement"`
	IntegrityRequirement       string `json:"integrity_requirement"`
	AvailabilityRequirement    string `json:"availability_requirement"`
}

type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Last string `json:"last"`
}
