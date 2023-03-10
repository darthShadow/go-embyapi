/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncSyncJobStatus string

// List of Sync.SyncJobStatus
const (
	QUEUED_SyncSyncJobStatus               SyncSyncJobStatus = "Queued"
	CONVERTING_SyncSyncJobStatus           SyncSyncJobStatus = "Converting"
	READY_TO_TRANSFER_SyncSyncJobStatus    SyncSyncJobStatus = "ReadyToTransfer"
	TRANSFERRING_SyncSyncJobStatus         SyncSyncJobStatus = "Transferring"
	COMPLETED_SyncSyncJobStatus            SyncSyncJobStatus = "Completed"
	COMPLETED_WITH_ERROR_SyncSyncJobStatus SyncSyncJobStatus = "CompletedWithError"
	FAILED_SyncSyncJobStatus               SyncSyncJobStatus = "Failed"
)
