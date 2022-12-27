/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UpdateUserEasyPassword struct {
	Id            string `json:"Id,omitempty"`
	NewPw         string `json:"NewPw,omitempty"`
	ResetPassword bool   `json:"ResetPassword,omitempty"`
}
