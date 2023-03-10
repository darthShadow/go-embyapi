/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SeriesDisplayOrder string

// List of SeriesDisplayOrder
const (
	AIRED_SeriesDisplayOrder    SeriesDisplayOrder = "Aired"
	DVD_SeriesDisplayOrder      SeriesDisplayOrder = "Dvd"
	ABSOLUTE_SeriesDisplayOrder SeriesDisplayOrder = "Absolute"
)
