/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyNotificationsNotificationCategoryInfo struct {
	Name   string                                  `json:"Name,omitempty"`
	Id     string                                  `json:"Id,omitempty"`
	Events []EmbyNotificationsNotificationTypeInfo `json:"Events,omitempty"`
}
