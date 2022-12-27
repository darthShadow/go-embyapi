/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaEncodingApiOnPlaybackProgress struct {
	PlaylistIndex  int32          `json:"PlaylistIndex,omitempty"`
	PlaylistLength int32          `json:"PlaylistLength,omitempty"`
	EventName      *ProgressEvent `json:"EventName,omitempty"`
}
