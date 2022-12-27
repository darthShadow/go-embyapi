/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaSourceType string

// List of MediaSourceType
const (
	DEFAULT__MediaSourceType    MediaSourceType = "Default"
	GROUPING_MediaSourceType    MediaSourceType = "Grouping"
	PLACEHOLDER_MediaSourceType MediaSourceType = "Placeholder"
)
