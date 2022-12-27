/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type CreateUserByName struct {
	Name            string                   `json:"Name,omitempty"`
	CopyFromUserId  string                   `json:"CopyFromUserId,omitempty"`
	UserCopyOptions []LibraryUserCopyOptions `json:"UserCopyOptions,omitempty"`
}
