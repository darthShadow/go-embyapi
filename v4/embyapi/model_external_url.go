/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ExternalUrl struct {
	Name string `json:"Name,omitempty"`
	Url  string `json:"Url,omitempty"`
}
