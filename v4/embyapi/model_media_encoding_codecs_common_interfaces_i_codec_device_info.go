/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaEncodingCodecsCommonInterfacesICodecDeviceInfo struct {
	Capabilities             *MediaEncodingCodecsCommonInterfacesICodecDeviceCapabilities `json:"Capabilities,omitempty"`
	Adapter                  int32                                                        `json:"Adapter,omitempty"`
	Name                     string                                                       `json:"Name,omitempty"`
	Desription               string                                                       `json:"Desription,omitempty"`
	Driver                   string                                                       `json:"Driver,omitempty"`
	DriverVersion            *Version                                                     `json:"DriverVersion,omitempty"`
	ApiVersion               *Version                                                     `json:"ApiVersion,omitempty"`
	VendorId                 int32                                                        `json:"VendorId,omitempty"`
	DeviceId                 int32                                                        `json:"DeviceId,omitempty"`
	DeviceIdentifier         string                                                       `json:"DeviceIdentifier,omitempty"`
	HardwareContextFramework *EmbyMediaModelEnumsSecondaryFrameworks                      `json:"HardwareContextFramework,omitempty"`
	DevPath                  string                                                       `json:"DevPath,omitempty"`
	DrmNode                  string                                                       `json:"DrmNode,omitempty"`
	VendorName               string                                                       `json:"VendorName,omitempty"`
	DeviceName               string                                                       `json:"DeviceName,omitempty"`
}
