/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type RokuMetadataApiThumbnailInfo struct {
	PositionTicks int64  `json:"PositionTicks,omitempty"`
	ImageTag      string `json:"ImageTag,omitempty"`
}
