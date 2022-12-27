/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ThemeMediaResult struct {
	OwnerId          int64         `json:"OwnerId,omitempty"`
	Items            []BaseItemDto `json:"Items,omitempty"`
	TotalRecordCount int32         `json:"TotalRecordCount,omitempty"`
}
