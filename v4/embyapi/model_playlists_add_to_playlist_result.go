/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PlaylistsAddToPlaylistResult struct {
	Id             string `json:"Id,omitempty"`
	ItemAddedCount int32  `json:"ItemAddedCount,omitempty"`
}
