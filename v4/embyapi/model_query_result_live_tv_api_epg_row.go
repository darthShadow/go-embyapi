/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultLiveTvApiEpgRow struct {
	Items            []LiveTvApiEpgRow `json:"Items,omitempty"`
	TotalRecordCount int32             `json:"TotalRecordCount,omitempty"`
}
