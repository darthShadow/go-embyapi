/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type NotificationsNotificationLevel string

// List of Notifications.NotificationLevel
const (
	NORMAL_NotificationsNotificationLevel  NotificationsNotificationLevel = "Normal"
	WARNING_NotificationsNotificationLevel NotificationsNotificationLevel = "Warning"
	ERROR__NotificationsNotificationLevel  NotificationsNotificationLevel = "Error"
)
