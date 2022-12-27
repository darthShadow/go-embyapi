/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyDlnaProfilesHttpHeaderInfo struct {
	Name  string                           `json:"Name,omitempty"`
	Value string                           `json:"Value,omitempty"`
	Match *EmbyDlnaProfilesHeaderMatchType `json:"Match,omitempty"`
}
