/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PlayMethod string

// List of PlayMethod
const (
	TRANSCODE_PlayMethod     PlayMethod = "Transcode"
	DIRECT_STREAM_PlayMethod PlayMethod = "DirectStream"
	DIRECT_PLAY_PlayMethod   PlayMethod = "DirectPlay"
)
