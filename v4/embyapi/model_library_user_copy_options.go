/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryUserCopyOptions string

// List of Library.UserCopyOptions
const (
	USER_POLICY_LibraryUserCopyOptions        LibraryUserCopyOptions = "UserPolicy"
	USER_CONFIGURATION_LibraryUserCopyOptions LibraryUserCopyOptions = "UserConfiguration"
)
