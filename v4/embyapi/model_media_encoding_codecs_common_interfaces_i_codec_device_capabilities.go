/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaEncodingCodecsCommonInterfacesICodecDeviceCapabilities struct {
	SupportsHwUpload             bool `json:"SupportsHwUpload,omitempty"`
	SupportsHwDownload           bool `json:"SupportsHwDownload,omitempty"`
	SupportsStandaloneDeviceInit bool `json:"SupportsStandaloneDeviceInit,omitempty"`
	Supports10BitProcessing      bool `json:"Supports10BitProcessing,omitempty"`
	SupportsNativeToneMapping    bool `json:"SupportsNativeToneMapping,omitempty"`
}
