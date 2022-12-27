/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UsersPinRedeemResult struct {
	Success    bool     `json:"Success,omitempty"`
	UsersReset []string `json:"UsersReset,omitempty"`
}
