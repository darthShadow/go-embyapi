/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConnectConnectAuthenticationExchangeResult struct {
	LocalUserId string `json:"LocalUserId,omitempty"`
	AccessToken string `json:"AccessToken,omitempty"`
}
