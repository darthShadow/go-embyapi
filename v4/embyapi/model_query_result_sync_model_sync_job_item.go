/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultSyncModelSyncJobItem struct {
	Items            []SyncModelSyncJobItem `json:"Items,omitempty"`
	TotalRecordCount int32                  `json:"TotalRecordCount,omitempty"`
}
