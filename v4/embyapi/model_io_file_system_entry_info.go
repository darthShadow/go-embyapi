/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type IoFileSystemEntryInfo struct {
	Name  string                 `json:"Name,omitempty"`
	Path  string                 `json:"Path,omitempty"`
	Type_ *IoFileSystemEntryType `json:"Type,omitempty"`
}
