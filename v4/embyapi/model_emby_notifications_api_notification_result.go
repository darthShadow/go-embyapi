/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyNotificationsApiNotificationResult struct {
	Notifications    []EmbyNotificationsApiNotification `json:"Notifications,omitempty"`
	TotalRecordCount int32                              `json:"TotalRecordCount,omitempty"`
}