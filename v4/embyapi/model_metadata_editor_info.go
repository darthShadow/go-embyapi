/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MetadataEditorInfo struct {
	ParentalRatingOptions []ParentalRating           `json:"ParentalRatingOptions,omitempty"`
	Countries             []GlobalizationCountryInfo `json:"Countries,omitempty"`
	Cultures              []GlobalizationCultureDto  `json:"Cultures,omitempty"`
	ExternalIdInfos       []ExternalIdInfo           `json:"ExternalIdInfos,omitempty"`
}
