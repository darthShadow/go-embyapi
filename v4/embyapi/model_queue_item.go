/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueueItem struct {
	Id             int64  `json:"Id,omitempty"`
	PlaylistItemId string `json:"PlaylistItemId,omitempty"`
}
