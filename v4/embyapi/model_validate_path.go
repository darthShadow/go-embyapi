/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ValidatePath struct {
	ValidateWriteable bool   `json:"ValidateWriteable,omitempty"`
	IsFile            bool   `json:"IsFile,omitempty"`
	Username          string `json:"Username,omitempty"`
	Password          string `json:"Password,omitempty"`
}
