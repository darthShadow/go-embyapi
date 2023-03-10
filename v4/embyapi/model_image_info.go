/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ImageInfo struct {
	ImageType  *ImageType `json:"ImageType,omitempty"`
	ImageIndex int32      `json:"ImageIndex,omitempty"`
	Path       string     `json:"Path,omitempty"`
	Filename   string     `json:"Filename,omitempty"`
	Height     int32      `json:"Height,omitempty"`
	Width      int32      `json:"Width,omitempty"`
	Size       int64      `json:"Size,omitempty"`
}
