/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UpdateUserPassword struct {
	Id            string `json:"Id,omitempty"`
	CurrentPw     string `json:"CurrentPw,omitempty"`
	NewPw         string `json:"NewPw,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}
