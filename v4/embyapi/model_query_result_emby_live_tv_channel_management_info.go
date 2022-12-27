/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultEmbyLiveTvChannelManagementInfo struct {
	Items            []EmbyLiveTvChannelManagementInfo `json:"Items,omitempty"`
	TotalRecordCount int32                             `json:"TotalRecordCount,omitempty"`
}
