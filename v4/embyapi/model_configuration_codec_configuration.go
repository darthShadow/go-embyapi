/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationCodecConfiguration struct {
	IsEnabled bool   `json:"IsEnabled,omitempty"`
	Priority  int32  `json:"Priority,omitempty"`
	CodecId   string `json:"CodecId,omitempty"`
}
