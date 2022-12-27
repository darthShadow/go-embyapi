/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaResponseProfile struct {
	Container  string                 `json:"Container,omitempty"`
	AudioCodec string                 `json:"AudioCodec,omitempty"`
	VideoCodec string                 `json:"VideoCodec,omitempty"`
	Type_      *DlnaDlnaProfileType   `json:"Type,omitempty"`
	OrgPn      string                 `json:"OrgPn,omitempty"`
	MimeType   string                 `json:"MimeType,omitempty"`
	Conditions []DlnaProfileCondition `json:"Conditions,omitempty"`
}
