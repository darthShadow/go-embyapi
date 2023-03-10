/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaSubtitleDeliveryMethod string

// List of Dlna.SubtitleDeliveryMethod
const (
	ENCODE_DlnaSubtitleDeliveryMethod          DlnaSubtitleDeliveryMethod = "Encode"
	EMBED_DlnaSubtitleDeliveryMethod           DlnaSubtitleDeliveryMethod = "Embed"
	EXTERNAL_DlnaSubtitleDeliveryMethod        DlnaSubtitleDeliveryMethod = "External"
	HLS_DlnaSubtitleDeliveryMethod             DlnaSubtitleDeliveryMethod = "Hls"
	VIDEO_SIDE_DATA_DlnaSubtitleDeliveryMethod DlnaSubtitleDeliveryMethod = "VideoSideData"
)
