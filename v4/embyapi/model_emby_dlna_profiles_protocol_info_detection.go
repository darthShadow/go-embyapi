/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyDlnaProfilesProtocolInfoDetection struct {
	EnabledForVideo  bool `json:"EnabledForVideo,omitempty"`
	EnabledForAudio  bool `json:"EnabledForAudio,omitempty"`
	EnabledForPhotos bool `json:"EnabledForPhotos,omitempty"`
}
