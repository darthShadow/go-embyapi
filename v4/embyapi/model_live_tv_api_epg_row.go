/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvApiEpgRow struct {
	Channel  *BaseItemDto  `json:"Channel,omitempty"`
	Programs []BaseItemDto `json:"Programs,omitempty"`
}
