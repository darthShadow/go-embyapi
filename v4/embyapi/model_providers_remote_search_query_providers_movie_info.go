/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ProvidersRemoteSearchQueryProvidersMovieInfo struct {
	SearchInfo               *ProvidersMovieInfo `json:"SearchInfo,omitempty"`
	ItemId                   int64               `json:"ItemId,omitempty"`
	SearchProviderName       string              `json:"SearchProviderName,omitempty"`
	Providers                []string            `json:"Providers,omitempty"`
	IncludeDisabledProviders bool                `json:"IncludeDisabledProviders,omitempty"`
}
