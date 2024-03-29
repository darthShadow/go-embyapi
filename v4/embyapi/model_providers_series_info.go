/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

import (
	"time"
)

type ProvidersSeriesInfo struct {
	EpisodeAirDate      time.Time                 `json:"EpisodeAirDate,omitempty"`
	DisplayOrder        *SeriesDisplayOrder       `json:"DisplayOrder,omitempty"`
	Name                string                    `json:"Name,omitempty"`
	MetadataLanguage    string                    `json:"MetadataLanguage,omitempty"`
	MetadataCountryCode string                    `json:"MetadataCountryCode,omitempty"`
	MetadataLanguages   []GlobalizationCultureDto `json:"MetadataLanguages,omitempty"`
	ProviderIds         *map[string]string        `json:"ProviderIds,omitempty"`
	Year                int32                     `json:"Year,omitempty"`
	IndexNumber         int32                     `json:"IndexNumber,omitempty"`
	ParentIndexNumber   int32                     `json:"ParentIndexNumber,omitempty"`
	PremiereDate        time.Time                 `json:"PremiereDate,omitempty"`
	IsAutomated         bool                      `json:"IsAutomated,omitempty"`
	EnableAdultMetadata bool                      `json:"EnableAdultMetadata,omitempty"`
}
