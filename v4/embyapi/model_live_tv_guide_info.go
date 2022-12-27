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

type LiveTvGuideInfo struct {
	StartDate time.Time `json:"StartDate,omitempty"`
	EndDate   time.Time `json:"EndDate,omitempty"`
}
