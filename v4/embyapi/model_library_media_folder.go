/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryMediaFolder struct {
	Name       string             `json:"Name,omitempty"`
	Id         string             `json:"Id,omitempty"`
	SubFolders []LibrarySubFolder `json:"SubFolders,omitempty"`
}
