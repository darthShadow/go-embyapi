/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultLiveTvSeriesTimerInfoDto struct {
	Items            []LiveTvSeriesTimerInfoDto `json:"Items,omitempty"`
	TotalRecordCount int32                      `json:"TotalRecordCount,omitempty"`
}
