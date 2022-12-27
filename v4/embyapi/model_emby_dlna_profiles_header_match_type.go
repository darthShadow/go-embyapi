/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyDlnaProfilesHeaderMatchType string

// List of Emby.Dlna.Profiles.HeaderMatchType
const (
	EQUALS_EmbyDlnaProfilesHeaderMatchType    EmbyDlnaProfilesHeaderMatchType = "Equals"
	REGEX_EmbyDlnaProfilesHeaderMatchType     EmbyDlnaProfilesHeaderMatchType = "Regex"
	SUBSTRING_EmbyDlnaProfilesHeaderMatchType EmbyDlnaProfilesHeaderMatchType = "Substring"
)
