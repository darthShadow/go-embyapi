/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationSegmentSkipMode string

// List of Configuration.SegmentSkipMode
const (
	SHOW_BUTTON_ConfigurationSegmentSkipMode ConfigurationSegmentSkipMode = "ShowButton"
	AUTO_SKIP_ConfigurationSegmentSkipMode   ConfigurationSegmentSkipMode = "AutoSkip"
	NONE_ConfigurationSegmentSkipMode        ConfigurationSegmentSkipMode = "None"
)
