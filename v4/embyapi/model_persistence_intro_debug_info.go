/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PersistenceIntroDebugInfo struct {
	Id    int64  `json:"Id,omitempty"`
	Path  string `json:"Path,omitempty"`
	Start int64  `json:"Start,omitempty"`
	End   int64  `json:"End,omitempty"`
}
