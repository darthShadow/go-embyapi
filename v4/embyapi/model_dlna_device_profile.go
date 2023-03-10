/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type DlnaDeviceProfile struct {
	Name                             string                   `json:"Name,omitempty"`
	Id                               string                   `json:"Id,omitempty"`
	SupportedMediaTypes              string                   `json:"SupportedMediaTypes,omitempty"`
	MaxStreamingBitrate              int64                    `json:"MaxStreamingBitrate,omitempty"`
	MusicStreamingTranscodingBitrate int32                    `json:"MusicStreamingTranscodingBitrate,omitempty"`
	MaxStaticMusicBitrate            int32                    `json:"MaxStaticMusicBitrate,omitempty"`
	DirectPlayProfiles               []DlnaDirectPlayProfile  `json:"DirectPlayProfiles,omitempty"`
	TranscodingProfiles              []DlnaTranscodingProfile `json:"TranscodingProfiles,omitempty"`
	ContainerProfiles                []DlnaContainerProfile   `json:"ContainerProfiles,omitempty"`
	CodecProfiles                    []DlnaCodecProfile       `json:"CodecProfiles,omitempty"`
	ResponseProfiles                 []DlnaResponseProfile    `json:"ResponseProfiles,omitempty"`
	SubtitleProfiles                 []DlnaSubtitleProfile    `json:"SubtitleProfiles,omitempty"`
}
