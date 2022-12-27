/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type NetEndPointInfo struct {
	IsLocal     bool `json:"IsLocal,omitempty"`
	IsInNetwork bool `json:"IsInNetwork,omitempty"`
}
