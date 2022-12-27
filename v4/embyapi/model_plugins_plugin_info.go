/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PluginsPluginInfo struct {
	Name                  string `json:"Name,omitempty"`
	Version               string `json:"Version,omitempty"`
	ConfigurationFileName string `json:"ConfigurationFileName,omitempty"`
	Description           string `json:"Description,omitempty"`
	Id                    string `json:"Id,omitempty"`
	ImageTag              string `json:"ImageTag,omitempty"`
}
