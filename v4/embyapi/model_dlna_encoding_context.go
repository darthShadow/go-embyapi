/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaEncodingContext string

// List of Dlna.EncodingContext
const (
	STREAMING_DlnaEncodingContext DlnaEncodingContext = "Streaming"
	STATIC_DlnaEncodingContext    DlnaEncodingContext = "Static"
)
