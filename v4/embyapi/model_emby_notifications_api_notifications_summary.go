/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyNotificationsApiNotificationsSummary struct {
	UnreadCount                int32                           `json:"UnreadCount,omitempty"`
	MaxUnreadNotificationLevel *NotificationsNotificationLevel `json:"MaxUnreadNotificationLevel,omitempty"`
}
