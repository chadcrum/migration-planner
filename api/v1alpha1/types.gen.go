// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"time"
)

// Defines values for InfraNetworksType.
const (
	Distributed InfraNetworksType = "distributed"
	Standard    InfraNetworksType = "standard"
)

// Defines values for SourceStatus.
const (
	SourceStatusError                     SourceStatus = "error"
	SourceStatusGatheringInitialInventory SourceStatus = "gathering-initial-inventory"
	SourceStatusNotConnected              SourceStatus = "not-connected"
	SourceStatusUpToDate                  SourceStatus = "up-to-date"
	SourceStatusWaitingForCredentials     SourceStatus = "waiting-for-credentials"
)

// Error defines model for Error.
type Error struct {
	// Message Error message
	Message string `json:"message"`
}

// Infra defines model for Infra.
type Infra struct {
	Datastores []struct {
		FreeCapacityGB  int    `json:"freeCapacityGB"`
		TotalCapacityGB int    `json:"totalCapacityGB"`
		Type            string `json:"type"`
	} `json:"datastores"`
	HostsPerCluster []int `json:"hostsPerCluster"`
	Networks        []struct {
		Name string            `json:"name"`
		Type InfraNetworksType `json:"type"`
	} `json:"networks"`
	TotalClusters int `json:"totalClusters"`
	TotalHosts    int `json:"totalHosts"`
}

// InfraNetworksType defines model for Infra.Networks.Type.
type InfraNetworksType string

// Inventory defines model for Inventory.
type Inventory struct {
	Infra Infra `json:"infra"`
	Vms   VMs   `json:"vms"`
}

// Source defines model for Source.
type Source struct {
	CreatedAt     time.Time    `json:"createdAt"`
	CredentialUrl string       `json:"credentialUrl"`
	Id            string       `json:"id"`
	Inventory     Inventory    `json:"inventory"`
	Name          string       `json:"name"`
	Status        SourceStatus `json:"status"`
	StatusInfo    string       `json:"statusInfo"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

// SourceStatus defines model for Source.Status.
type SourceStatus string

// SourceCreate defines model for SourceCreate.
type SourceCreate struct {
	Name string `json:"name"`
}

// SourceList defines model for SourceList.
type SourceList = []Source

// Status Status is a return value for calls that don't return other objects.
type Status struct {
	// Message A human-readable description of the status of this operation.
	Message *string `json:"message,omitempty"`

	// Reason A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason *string `json:"reason,omitempty"`

	// Status Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status *string `json:"status,omitempty"`
}

// VMResourceBreakdown defines model for VMResourceBreakdown.
type VMResourceBreakdown struct {
	Histogram struct {
		Data     []int `json:"data"`
		MinValue int   `json:"minValue"`
		Step     int   `json:"step"`
	} `json:"histogram"`
	Total                          int  `json:"total"`
	TotalForMigratable             *int `json:"totalForMigratable,omitempty"`
	TotalForMigratableWithWarnings *int `json:"totalForMigratableWithWarnings,omitempty"`
	TotalForNotMigratable          *int `json:"totalForNotMigratable,omitempty"`
}

// VMs defines model for VMs.
type VMs struct {
	CpuCores                    VMResourceBreakdown `json:"cpuCores"`
	DiskGB                      VMResourceBreakdown `json:"diskGB"`
	MigrationWarnings           map[string]int      `json:"migrationWarnings"`
	NotMigratableReasons        map[string]int      `json:"notMigratableReasons"`
	Os                          map[string]int      `json:"os"`
	RamGB                       VMResourceBreakdown `json:"ramGB"`
	Total                       int                 `json:"total"`
	TotalMigratable             int                 `json:"totalMigratable"`
	TotalMigratableWithWarnings *int                `json:"totalMigratableWithWarnings,omitempty"`
}

// CreateSourceJSONRequestBody defines body for CreateSource for application/json ContentType.
type CreateSourceJSONRequestBody = SourceCreate
