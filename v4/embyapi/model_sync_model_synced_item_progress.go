/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncedItemProgress struct {
	Progress float64                     `json:"Progress,omitempty"`
	Status   *SyncModelSyncJobItemStatus `json:"Status,omitempty"`
}
