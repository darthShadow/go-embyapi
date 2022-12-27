/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PlaystateRequest struct {
	Command           *PlaystateCommand `json:"Command,omitempty"`
	SeekPositionTicks int64             `json:"SeekPositionTicks,omitempty"`
	ControllingUserId string            `json:"ControllingUserId,omitempty"`
}
