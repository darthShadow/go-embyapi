/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type GlobalizationCultureDto struct {
	Name                        string   `json:"Name,omitempty"`
	DisplayName                 string   `json:"DisplayName,omitempty"`
	TwoLetterISOLanguageName    string   `json:"TwoLetterISOLanguageName,omitempty"`
	ThreeLetterISOLanguageName  string   `json:"ThreeLetterISOLanguageName,omitempty"`
	ThreeLetterISOLanguageNames []string `json:"ThreeLetterISOLanguageNames,omitempty"`
	TwoLetterISOLanguageNames   []string `json:"TwoLetterISOLanguageNames,omitempty"`
}
