/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyWebGenericUiApiEndpointsRunUiCommand struct {
	PageId       string `json:"PageId,omitempty"`
	CommandId    string `json:"CommandId,omitempty"`
	Data         string `json:"Data,omitempty"`
	ItemId       string `json:"ItemId,omitempty"`
	ClientLocale string `json:"ClientLocale,omitempty"`
}
