/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncJobItemStatus string

// List of Sync.Model.SyncJobItemStatus
const (
	QUEUED_SyncModelSyncJobItemStatus            SyncModelSyncJobItemStatus = "Queued"
	CONVERTING_SyncModelSyncJobItemStatus        SyncModelSyncJobItemStatus = "Converting"
	READY_TO_TRANSFER_SyncModelSyncJobItemStatus SyncModelSyncJobItemStatus = "ReadyToTransfer"
	TRANSFERRING_SyncModelSyncJobItemStatus      SyncModelSyncJobItemStatus = "Transferring"
	SYNCED_SyncModelSyncJobItemStatus            SyncModelSyncJobItemStatus = "Synced"
	FAILED_SyncModelSyncJobItemStatus            SyncModelSyncJobItemStatus = "Failed"
)
