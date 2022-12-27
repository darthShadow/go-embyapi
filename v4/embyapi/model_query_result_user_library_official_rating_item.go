/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultUserLibraryOfficialRatingItem struct {
	Items            []UserLibraryOfficialRatingItem `json:"Items,omitempty"`
	TotalRecordCount int32                           `json:"TotalRecordCount,omitempty"`
}
