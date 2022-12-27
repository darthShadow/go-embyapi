/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationTypeOptions struct {
	Type_                string                     `json:"Type,omitempty"`
	MetadataFetchers     []string                   `json:"MetadataFetchers,omitempty"`
	MetadataFetcherOrder []string                   `json:"MetadataFetcherOrder,omitempty"`
	ImageFetchers        []string                   `json:"ImageFetchers,omitempty"`
	ImageFetcherOrder    []string                   `json:"ImageFetcherOrder,omitempty"`
	ImageOptions         []ConfigurationImageOption `json:"ImageOptions,omitempty"`
}
