/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultLiveTvTimerInfoDto struct {
	Items            []LiveTvTimerInfoDto `json:"Items,omitempty"`
	TotalRecordCount int32                `json:"TotalRecordCount,omitempty"`
}
