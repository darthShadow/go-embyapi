/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvLiveTvInfo struct {
	IsEnabled    bool     `json:"IsEnabled,omitempty"`
	EnabledUsers []string `json:"EnabledUsers,omitempty"`
}
