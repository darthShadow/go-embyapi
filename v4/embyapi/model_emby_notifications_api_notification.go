/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

import (
	"time"
)

type EmbyNotificationsApiNotification struct {
	Id          string                          `json:"Id,omitempty"`
	UserId      string                          `json:"UserId,omitempty"`
	Date        time.Time                       `json:"Date,omitempty"`
	IsRead      bool                            `json:"IsRead,omitempty"`
	Name        string                          `json:"Name,omitempty"`
	Description string                          `json:"Description,omitempty"`
	Url         string                          `json:"Url,omitempty"`
	Level       *NotificationsNotificationLevel `json:"Level,omitempty"`
}
