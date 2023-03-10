/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type GeneralCommand struct {
	Name              string            `json:"Name,omitempty"`
	ControllingUserId string            `json:"ControllingUserId,omitempty"`
	Arguments         map[string]string `json:"Arguments,omitempty"`
}
