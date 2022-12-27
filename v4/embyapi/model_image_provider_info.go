/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ImageProviderInfo struct {
	Name            string      `json:"Name,omitempty"`
	SupportedImages []ImageType `json:"SupportedImages,omitempty"`
}
