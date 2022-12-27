/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UserLibraryAddTags struct {
	Tags []NameIdPair `json:"Tags,omitempty"`
}
