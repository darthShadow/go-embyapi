/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type NameLongIdPair struct {
	Name string `json:"Name,omitempty"`
	Id   int64  `json:"Id,omitempty"`
}
