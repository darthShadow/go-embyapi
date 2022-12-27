/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type VirtualFolderInfo struct {
	Name               string                       `json:"Name,omitempty"`
	Locations          []string                     `json:"Locations,omitempty"`
	CollectionType     string                       `json:"CollectionType,omitempty"`
	LibraryOptions     *ConfigurationLibraryOptions `json:"LibraryOptions,omitempty"`
	ItemId             string                       `json:"ItemId,omitempty"`
	PrimaryImageItemId string                       `json:"PrimaryImageItemId,omitempty"`
	RefreshProgress    float64                      `json:"RefreshProgress,omitempty"`
	RefreshStatus      string                       `json:"RefreshStatus,omitempty"`
}