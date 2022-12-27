/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryRenameVirtualFolder struct {
	Id      string `json:"Id,omitempty"`
	NewName string `json:"NewName,omitempty"`
}
