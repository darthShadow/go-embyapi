/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyMediaModelTypesProfileLevelInformation struct {
	Profile *EmbyMediaModelTypesProfileInformation `json:"Profile,omitempty"`
	Level   *EmbyMediaModelTypesLevelInformation   `json:"Level,omitempty"`
}
