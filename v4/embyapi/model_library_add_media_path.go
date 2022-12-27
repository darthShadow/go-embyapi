/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryAddMediaPath struct {
	Id             string                      `json:"Id,omitempty"`
	Path           string                      `json:"Path,omitempty"`
	PathInfo       *ConfigurationMediaPathInfo `json:"PathInfo,omitempty"`
	RefreshLibrary bool                        `json:"RefreshLibrary,omitempty"`
}
