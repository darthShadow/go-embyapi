/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConnectUserLinkResult struct {
	IsPending           bool   `json:"IsPending,omitempty"`
	IsNewUserInvitation bool   `json:"IsNewUserInvitation,omitempty"`
	GuestDisplayName    string `json:"GuestDisplayName,omitempty"`
}
