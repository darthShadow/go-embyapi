/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryUpdateMediaPath struct {
	Id       string                      `json:"Id,omitempty"`
	PathInfo *ConfigurationMediaPathInfo `json:"PathInfo,omitempty"`
}
