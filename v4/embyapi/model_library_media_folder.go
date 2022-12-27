/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryMediaFolder struct {
	Name       string             `json:"Name,omitempty"`
	Id         string             `json:"Id,omitempty"`
	Guid       string             `json:"Guid,omitempty"`
	SubFolders []LibrarySubFolder `json:"SubFolders,omitempty"`
}
