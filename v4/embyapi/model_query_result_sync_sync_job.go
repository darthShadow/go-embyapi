/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultSyncSyncJob struct {
	Items            []SyncSyncJob `json:"Items,omitempty"`
	TotalRecordCount int32         `json:"TotalRecordCount,omitempty"`
}
