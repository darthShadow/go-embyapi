/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryLibraryTypeOptions struct {
	Type_               string                     `json:"Type,omitempty"`
	MetadataFetchers    []LibraryLibraryOptionInfo `json:"MetadataFetchers,omitempty"`
	ImageFetchers       []LibraryLibraryOptionInfo `json:"ImageFetchers,omitempty"`
	SupportedImageTypes []ImageType                `json:"SupportedImageTypes,omitempty"`
	DefaultImageOptions []ConfigurationImageOption `json:"DefaultImageOptions,omitempty"`
}
