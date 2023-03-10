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

type UpdatesPackageInfo struct {
	Id               string                      `json:"id,omitempty"`
	Name             string                      `json:"name,omitempty"`
	ShortDescription string                      `json:"shortDescription,omitempty"`
	Overview         string                      `json:"overview,omitempty"`
	IsPremium        bool                        `json:"isPremium,omitempty"`
	Adult            bool                        `json:"adult,omitempty"`
	RichDescUrl      string                      `json:"richDescUrl,omitempty"`
	ThumbImage       string                      `json:"thumbImage,omitempty"`
	PreviewImage     string                      `json:"previewImage,omitempty"`
	Type_            string                      `json:"type,omitempty"`
	TargetFilename   string                      `json:"targetFilename,omitempty"`
	Owner            string                      `json:"owner,omitempty"`
	Category         string                      `json:"category,omitempty"`
	TileColor        string                      `json:"tileColor,omitempty"`
	FeatureId        string                      `json:"featureId,omitempty"`
	Price            float32                     `json:"price,omitempty"`
	TargetSystem     *UpdatesPackageTargetSystem `json:"targetSystem,omitempty"`
	Guid             string                      `json:"guid,omitempty"`
	IsRegistered     bool                        `json:"isRegistered,omitempty"`
	ExpDate          time.Time                   `json:"expDate,omitempty"`
	Versions         []UpdatesPackageVersionInfo `json:"versions,omitempty"`
	EnableInAppStore bool                        `json:"enableInAppStore,omitempty"`
	Installs         int32                       `json:"installs,omitempty"`
}
