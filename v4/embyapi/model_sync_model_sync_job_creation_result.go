/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncJobCreationResult struct {
	Job      *SyncSyncJob           `json:"Job,omitempty"`
	JobItems []SyncModelSyncJobItem `json:"JobItems,omitempty"`
}
