/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type AllThemeMediaResult struct {
	ThemeVideosResult     *ThemeMediaResult `json:"ThemeVideosResult,omitempty"`
	ThemeSongsResult      *ThemeMediaResult `json:"ThemeSongsResult,omitempty"`
	SoundtrackSongsResult *ThemeMediaResult `json:"SoundtrackSongsResult,omitempty"`
}
