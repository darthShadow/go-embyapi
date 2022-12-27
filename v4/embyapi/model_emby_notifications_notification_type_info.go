/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyNotificationsNotificationTypeInfo struct {
	Name         string `json:"Name,omitempty"`
	Id           string `json:"Id,omitempty"`
	CategoryName string `json:"CategoryName,omitempty"`
	CategoryId   string `json:"CategoryId,omitempty"`
}
