/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncQualityOption struct {
	Name              string `json:"Name,omitempty"`
	Description       string `json:"Description,omitempty"`
	Id                string `json:"Id,omitempty"`
	IsDefault         bool   `json:"IsDefault,omitempty"`
	IsOriginalQuality bool   `json:"IsOriginalQuality,omitempty"`
}
