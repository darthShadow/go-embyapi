/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvApiSetChannelDisabled struct {
	Id           string `json:"Id,omitempty"`
	ManagementId string `json:"ManagementId,omitempty"`
	Disabled     bool   `json:"Disabled,omitempty"`
}
