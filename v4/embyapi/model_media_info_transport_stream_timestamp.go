/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaInfoTransportStreamTimestamp string

// List of MediaInfo.TransportStreamTimestamp
const (
	NONE_MediaInfoTransportStreamTimestamp  MediaInfoTransportStreamTimestamp = "None"
	ZERO_MediaInfoTransportStreamTimestamp  MediaInfoTransportStreamTimestamp = "Zero"
	VALID_MediaInfoTransportStreamTimestamp MediaInfoTransportStreamTimestamp = "Valid"
)
