/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvApiSetChannelSortIndex struct {
	Id           string `json:"Id,omitempty"`
	ManagementId string `json:"ManagementId,omitempty"`
	NewIndex     int32  `json:"NewIndex,omitempty"`
}
