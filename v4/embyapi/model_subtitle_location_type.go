/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SubtitleLocationType string

// List of SubtitleLocationType
const (
	INTERNAL_STREAM_SubtitleLocationType SubtitleLocationType = "InternalStream"
	VIDEO_SIDE_DATA_SubtitleLocationType SubtitleLocationType = "VideoSideData"
)
