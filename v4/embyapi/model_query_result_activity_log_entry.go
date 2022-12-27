/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultActivityLogEntry struct {
	Items            []ActivityLogEntry `json:"Items,omitempty"`
	TotalRecordCount int32              `json:"TotalRecordCount,omitempty"`
}
