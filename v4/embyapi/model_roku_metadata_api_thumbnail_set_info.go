/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type RokuMetadataApiThumbnailSetInfo struct {
	AspectRatio float64                        `json:"AspectRatio,omitempty"`
	Thumbnails  []RokuMetadataApiThumbnailInfo `json:"Thumbnails,omitempty"`
}
