/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationMetadataFeatures string

// List of Configuration.MetadataFeatures
const (
	COLLECTIONS_ConfigurationMetadataFeatures    ConfigurationMetadataFeatures = "Collections"
	ADULT_ConfigurationMetadataFeatures          ConfigurationMetadataFeatures = "Adult"
	REQUIRED_SETUP_ConfigurationMetadataFeatures ConfigurationMetadataFeatures = "RequiredSetup"
)
