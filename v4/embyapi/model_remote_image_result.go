/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type RemoteImageResult struct {
	Images           []RemoteImageInfo `json:"Images,omitempty"`
	TotalRecordCount int32             `json:"TotalRecordCount,omitempty"`
	Providers        []string          `json:"Providers,omitempty"`
}
