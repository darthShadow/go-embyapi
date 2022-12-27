/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ParentalRating struct {
	Name  string `json:"Name,omitempty"`
	Value int32  `json:"Value,omitempty"`
}
