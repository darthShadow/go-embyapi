/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultDevicesDeviceInfo struct {
	Items            []DevicesDeviceInfo `json:"Items,omitempty"`
	TotalRecordCount int32               `json:"TotalRecordCount,omitempty"`
}
