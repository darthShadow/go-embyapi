/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaDirectPlayProfile struct {
	Container  string               `json:"Container,omitempty"`
	AudioCodec string               `json:"AudioCodec,omitempty"`
	VideoCodec string               `json:"VideoCodec,omitempty"`
	Type_      *DlnaDlnaProfileType `json:"Type,omitempty"`
}
