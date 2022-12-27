/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultString struct {
	Items            []string `json:"Items,omitempty"`
	TotalRecordCount int32    `json:"TotalRecordCount,omitempty"`
}
