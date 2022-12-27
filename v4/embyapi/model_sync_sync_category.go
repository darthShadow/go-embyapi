/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncSyncCategory string

// List of Sync.SyncCategory
const (
	LATEST_SyncSyncCategory  SyncSyncCategory = "Latest"
	NEXT_UP_SyncSyncCategory SyncSyncCategory = "NextUp"
	RESUME_SyncSyncCategory  SyncSyncCategory = "Resume"
)
