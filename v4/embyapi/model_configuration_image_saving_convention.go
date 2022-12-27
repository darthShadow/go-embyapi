/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationImageSavingConvention string

// List of Configuration.ImageSavingConvention
const (
	LEGACY_ConfigurationImageSavingConvention     ConfigurationImageSavingConvention = "Legacy"
	COMPATIBLE_ConfigurationImageSavingConvention ConfigurationImageSavingConvention = "Compatible"
)
