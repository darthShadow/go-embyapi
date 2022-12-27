/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryRemoveVirtualFolder struct {
	Id             string `json:"Id,omitempty"`
	RefreshLibrary bool   `json:"RefreshLibrary,omitempty"`
}
