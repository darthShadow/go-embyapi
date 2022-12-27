/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryUpdateLibraryOptions struct {
	Id             string                       `json:"Id,omitempty"`
	LibraryOptions *ConfigurationLibraryOptions `json:"LibraryOptions,omitempty"`
}
