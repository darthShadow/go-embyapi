/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaDlnaProfileType string

// List of Dlna.DlnaProfileType
const (
	AUDIO_DlnaDlnaProfileType DlnaDlnaProfileType = "Audio"
	VIDEO_DlnaDlnaProfileType DlnaDlnaProfileType = "Video"
	PHOTO_DlnaDlnaProfileType DlnaDlnaProfileType = "Photo"
)
