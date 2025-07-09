package hackerone

type ScopeResponse struct {
	Data  []Data `json:"data"`
	Links Links  `json:"links"`
}

type Scope struct {
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Program struct {
	ID           string `json:"id"`
	Handle       string `json:"handle"`
	HandleApiUrl string
}

type Attributes struct {
	ID                         string `json:"id"`
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
