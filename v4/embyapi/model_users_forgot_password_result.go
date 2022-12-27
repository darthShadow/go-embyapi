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

type UsersForgotPasswordResult struct {
	Action            *UsersForgotPasswordAction `json:"Action,omitempty"`
	PinFile           string                     `json:"PinFile,omitempty"`
	PinExpirationDate time.Time                  `json:"PinExpirationDate,omitempty"`
}
