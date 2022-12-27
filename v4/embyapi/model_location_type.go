/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LocationType string

// List of LocationType
const (
	FILE_SYSTEM_LocationType LocationType = "FileSystem"
	VIRTUAL_LocationType     LocationType = "Virtual"
)
