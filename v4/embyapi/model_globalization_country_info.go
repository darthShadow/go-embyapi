/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type GlobalizationCountryInfo struct {
	Name                     string `json:"Name,omitempty"`
	DisplayName              string `json:"DisplayName,omitempty"`
	EnglishName              string `json:"EnglishName,omitempty"`
	TwoLetterISORegionName   string `json:"TwoLetterISORegionName,omitempty"`
	ThreeLetterISORegionName string `json:"ThreeLetterISORegionName,omitempty"`
}
