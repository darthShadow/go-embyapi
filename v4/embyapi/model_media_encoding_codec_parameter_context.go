/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaEncodingCodecParameterContext string

// List of MediaEncoding.CodecParameterContext
const (
	PLAYBACK_MediaEncodingCodecParameterContext   MediaEncodingCodecParameterContext = "Playback"
	CONVERSION_MediaEncodingCodecParameterContext MediaEncodingCodecParameterContext = "Conversion"
)
