/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationImageOption struct {
	Type_    *ImageType `json:"Type,omitempty"`
	Limit    int32      `json:"Limit,omitempty"`
	MinWidth int32      `json:"MinWidth,omitempty"`
}
