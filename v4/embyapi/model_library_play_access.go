/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryPlayAccess string

// List of Library.PlayAccess
const (
	FULL_LibraryPlayAccess LibraryPlayAccess = "Full"
	NONE_LibraryPlayAccess LibraryPlayAccess = "None"
)
