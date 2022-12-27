/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConnectUserLinkType string

// List of Connect.UserLinkType
const (
	LINKED_USER_ConnectUserLinkType ConnectUserLinkType = "LinkedUser"
	GUEST_ConnectUserLinkType       ConnectUserLinkType = "Guest"
)
