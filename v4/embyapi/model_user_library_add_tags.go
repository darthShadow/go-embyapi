/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UserLibraryAddTags struct {
	Tags []NameIdPair `json:"Tags,omitempty"`
}
