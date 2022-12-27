/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvApiSetChannelMapping struct {
	TunerChannelId    string `json:"TunerChannelId,omitempty"`
	ProviderChannelId string `json:"ProviderChannelId,omitempty"`
}
