/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaSubtitleProfile struct {
	Format    string                      `json:"Format,omitempty"`
	Method    *DlnaSubtitleDeliveryMethod `json:"Method,omitempty"`
	DidlMode  string                      `json:"DidlMode,omitempty"`
	Language  string                      `json:"Language,omitempty"`
	Container string                      `json:"Container,omitempty"`
	Protocol  string                      `json:"Protocol,omitempty"`
}
