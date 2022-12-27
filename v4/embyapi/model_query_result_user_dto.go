/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultUserDto struct {
	Items            []UserDto `json:"Items,omitempty"`
	TotalRecordCount int32     `json:"TotalRecordCount,omitempty"`
}
