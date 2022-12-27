/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelItemFileType string

// List of Sync.Model.ItemFileType
const (
	MEDIA_SyncModelItemFileType     SyncModelItemFileType = "Media"
	IMAGE_SyncModelItemFileType     SyncModelItemFileType = "Image"
	SUBTITLES_SyncModelItemFileType SyncModelItemFileType = "Subtitles"
)
