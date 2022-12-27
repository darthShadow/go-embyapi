/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type AuthenticateUser struct {
	Pw string `json:"Pw,omitempty"`
}
