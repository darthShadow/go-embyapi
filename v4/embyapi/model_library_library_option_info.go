/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryLibraryOptionInfo struct {
	Name           string                          `json:"Name,omitempty"`
	SetupUrl       string                          `json:"SetupUrl,omitempty"`
	DefaultEnabled bool                            `json:"DefaultEnabled,omitempty"`
	Features       []ConfigurationMetadataFeatures `json:"Features,omitempty"`
}
