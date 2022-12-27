/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SessionUserInfo struct {
	UserId         string `json:"UserId,omitempty"`
	UserName       string `json:"UserName,omitempty"`
	UserInternalId int64  `json:"UserInternalId,omitempty"`
}
