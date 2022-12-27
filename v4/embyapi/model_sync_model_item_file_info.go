/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type SyncModelItemFileInfo struct {
	Type_     *SyncModelItemFileType `json:"Type,omitempty"`
	Name      string                 `json:"Name,omitempty"`
	Path      string                 `json:"Path,omitempty"`
	ImageType *ImageType             `json:"ImageType,omitempty"`
	Index     int32                  `json:"Index,omitempty"`
}
