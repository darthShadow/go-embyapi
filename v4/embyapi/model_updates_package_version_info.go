/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

import (
	"time"
)

type UpdatesPackageVersionInfo struct {
	Name               string                      `json:"name,omitempty"`
	Guid               string                      `json:"guid,omitempty"`
	VersionStr         string                      `json:"versionStr,omitempty"`
	Classification     *UpdatesPackageVersionClass `json:"classification,omitempty"`
	Description        string                      `json:"description,omitempty"`
	RequiredVersionStr string                      `json:"requiredVersionStr,omitempty"`
	SourceUrl          string                      `json:"sourceUrl,omitempty"`
	Checksum           string                      `json:"checksum,omitempty"`
	TargetFilename     string                      `json:"targetFilename,omitempty"`
	InfoUrl            string                      `json:"infoUrl,omitempty"`
	Runtimes           string                      `json:"runtimes,omitempty"`
	Timestamp          time.Time                   `json:"timestamp,omitempty"`
}
