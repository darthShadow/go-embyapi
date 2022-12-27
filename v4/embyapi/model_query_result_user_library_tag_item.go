/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultUserLibraryTagItem struct {
	Items            []UserLibraryTagItem `json:"Items,omitempty"`
	TotalRecordCount int32                `json:"TotalRecordCount,omitempty"`
}
