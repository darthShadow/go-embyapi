/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type ConfigurationMediaPathInfo struct {
	Path        string `json:"Path,omitempty"`
	NetworkPath string `json:"NetworkPath,omitempty"`
	Username    string `json:"Username,omitempty"`
	Password    string `json:"Password,omitempty"`
}
