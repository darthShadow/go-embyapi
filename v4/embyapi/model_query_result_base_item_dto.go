/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultBaseItemDto struct {
	Items            []BaseItemDto `json:"Items,omitempty"`
	TotalRecordCount int32         `json:"TotalRecordCount,omitempty"`
}
