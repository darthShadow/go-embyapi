/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyMediaModelTypesResolutionWithRate struct {
	Width      int32                          `json:"Width,omitempty"`
	Height     int32                          `json:"Height,omitempty"`
	FrameRate  float64                        `json:"FrameRate,omitempty"`
	Resolution *EmbyMediaModelTypesResolution `json:"Resolution,omitempty"`
}
