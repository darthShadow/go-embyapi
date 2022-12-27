/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DevicesContentUploadHistory struct {
	DeviceId      string                 `json:"DeviceId,omitempty"`
	FilesUploaded []DevicesLocalFileInfo `json:"FilesUploaded,omitempty"`
}
