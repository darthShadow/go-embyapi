/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelSyncDataResponse struct {
	ItemIdsToRemove []string `json:"ItemIdsToRemove,omitempty"`
}
