/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationUserConfiguration struct {
	AudioLanguagePreference    string                             `json:"AudioLanguagePreference,omitempty"`
	PlayDefaultAudioTrack      bool                               `json:"PlayDefaultAudioTrack,omitempty"`
	SubtitleLanguagePreference string                             `json:"SubtitleLanguagePreference,omitempty"`
	DisplayMissingEpisodes     bool                               `json:"DisplayMissingEpisodes,omitempty"`
	SubtitleMode               *ConfigurationSubtitlePlaybackMode `json:"SubtitleMode,omitempty"`
	OrderedViews               []string                           `json:"OrderedViews,omitempty"`
	LatestItemsExcludes        []string                           `json:"LatestItemsExcludes,omitempty"`
	MyMediaExcludes            []string                           `json:"MyMediaExcludes,omitempty"`
	HidePlayedInLatest         bool                               `json:"HidePlayedInLatest,omitempty"`
	RememberAudioSelections    bool                               `json:"RememberAudioSelections,omitempty"`
	RememberSubtitleSelections bool                               `json:"RememberSubtitleSelections,omitempty"`
	EnableNextEpisodeAutoPlay  bool                               `json:"EnableNextEpisodeAutoPlay,omitempty"`
	ResumeRewindSeconds        int32                              `json:"ResumeRewindSeconds,omitempty"`
	IntroSkipMode              *ConfigurationSegmentSkipMode      `json:"IntroSkipMode,omitempty"`
	EnableLocalPassword        bool                               `json:"EnableLocalPassword,omitempty"`
}
