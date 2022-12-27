/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultLogFile struct {
	Items            []LogFile `json:"Items,omitempty"`
	TotalRecordCount int32     `json:"TotalRecordCount,omitempty"`
}
