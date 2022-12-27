/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type QueryResultVirtualFolderInfo struct {
	Items            []VirtualFolderInfo `json:"Items,omitempty"`
	TotalRecordCount int32               `json:"TotalRecordCount,omitempty"`
}
