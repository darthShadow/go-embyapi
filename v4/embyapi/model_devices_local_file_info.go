/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DevicesLocalFileInfo struct {
	Name     string `json:"Name,omitempty"`
	Id       string `json:"Id,omitempty"`
	Album    string `json:"Album,omitempty"`
	MimeType string `json:"MimeType,omitempty"`
}
