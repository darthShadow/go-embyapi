/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncDataResponse struct {
	ItemIdsToRemove []string `json:"ItemIdsToRemove,omitempty"`
}
