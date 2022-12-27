/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibrarySubFolder struct {
	Name string `json:"Name,omitempty"`
	Id   string `json:"Id,omitempty"`
	Path string `json:"Path,omitempty"`
}
