/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryRemoveMediaPath struct {
	Id             string `json:"Id,omitempty"`
	Path           string `json:"Path,omitempty"`
	RefreshLibrary bool   `json:"RefreshLibrary,omitempty"`
}
