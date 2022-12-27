/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaTranscodeSeekInfo string

// List of Dlna.TranscodeSeekInfo
const (
	AUTO_DlnaTranscodeSeekInfo  DlnaTranscodeSeekInfo = "Auto"
	BYTES_DlnaTranscodeSeekInfo DlnaTranscodeSeekInfo = "Bytes"
)
