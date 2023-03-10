/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

import (
	"time"
)

type UsersUserAction struct {
	Id            string               `json:"Id,omitempty"`
	ServerId      string               `json:"ServerId,omitempty"`
	UserId        string               `json:"UserId,omitempty"`
	ItemId        string               `json:"ItemId,omitempty"`
	Type_         *UsersUserActionType `json:"Type,omitempty"`
	Date          time.Time            `json:"Date,omitempty"`
	PositionTicks int64                `json:"PositionTicks,omitempty"`
}
