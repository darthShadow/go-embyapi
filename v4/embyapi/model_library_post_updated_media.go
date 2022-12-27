/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryPostUpdatedMedia struct {
	Updates []LibraryMediaUpdateInfo `json:"Updates,omitempty"`
}
