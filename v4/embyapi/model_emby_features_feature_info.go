/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyFeaturesFeatureInfo struct {
	Name        string                   `json:"Name,omitempty"`
	Id          string                   `json:"Id,omitempty"`
	FeatureType *EmbyFeaturesFeatureType `json:"FeatureType,omitempty"`
}
