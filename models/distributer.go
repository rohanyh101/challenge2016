package models

type Distributor struct {
	DistributorID     string   `json:"distributor_id"`
	Name              string   `json:"name"`
	Includes          []string `json:"includes,omitempty"`
	Excludes          []string `json:"excludes,omitempty"`
	Level             int      `json:"level,omitempty"`
	ParentDistributor string   `json:"parent_distributor_id,omitempty"`
}

// commands used to interact with webapi
type NewDistributorCmd struct {
	Name                string   `json:"name"`
	Includes            []string `json:"includes,omitempty"`
	Excludes            []string `json:"excludes,omitempty"`
	ParentDistributorID string   `json:"parent_distributor_id"`
}

type CheckDistributorPermissionCmd struct {
	Locations []string `json:"locations"`
}

type CheckDistributorPermissionResponse struct {
	DistributorID   string          `json:"distributor_id"`
	DistributorName string          `json:"distributor_name"`
	PermissionMap   map[string]bool `json:"permission"`
}

type UpdateDistributorCmd struct {
	Name                string   `json:"name"`
	Includes            []string `json:"includes,omitempty"`
	Excludes            []string `json:"excludes,omitempty"`
	ParentDistributorID string   `json:"parent_distributor_id"`
}
