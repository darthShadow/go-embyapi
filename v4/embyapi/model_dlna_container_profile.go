/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaContainerProfile struct {
	Type_      *DlnaDlnaProfileType   `json:"Type,omitempty"`
	Conditions []DlnaProfileCondition `json:"Conditions,omitempty"`
	Container  string                 `json:"Container,omitempty"`
}
