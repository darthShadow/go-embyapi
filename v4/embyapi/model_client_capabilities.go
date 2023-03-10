/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ClientCapabilities struct {
	PlayableMediaTypes   []string           `json:"PlayableMediaTypes,omitempty"`
	SupportedCommands    []string           `json:"SupportedCommands,omitempty"`
	SupportsMediaControl bool               `json:"SupportsMediaControl,omitempty"`
	PushToken            string             `json:"PushToken,omitempty"`
	PushTokenType        string             `json:"PushTokenType,omitempty"`
	SupportsSync         bool               `json:"SupportsSync,omitempty"`
	DeviceProfile        *DlnaDeviceProfile `json:"DeviceProfile,omitempty"`
	IconUrl              string             `json:"IconUrl,omitempty"`
	AppId                string             `json:"AppId,omitempty"`
}
