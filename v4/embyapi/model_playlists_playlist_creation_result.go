/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PlaylistsPlaylistCreationResult struct {
	Id             string `json:"Id,omitempty"`
	Name           string `json:"Name,omitempty"`
	ItemAddedCount int32  `json:"ItemAddedCount,omitempty"`
}
