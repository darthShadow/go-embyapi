/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaInfoPlaybackInfoResponse struct {
	MediaSources  []MediaSourceInfo      `json:"MediaSources,omitempty"`
	PlaySessionId string                 `json:"PlaySessionId,omitempty"`
	ErrorCode     *DlnaPlaybackErrorCode `json:"ErrorCode,omitempty"`
}
