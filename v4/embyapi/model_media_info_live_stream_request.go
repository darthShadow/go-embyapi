/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaInfoLiveStreamRequest struct {
	OpenToken                      string                   `json:"OpenToken,omitempty"`
	UserId                         string                   `json:"UserId,omitempty"`
	PlaySessionId                  string                   `json:"PlaySessionId,omitempty"`
	MaxStreamingBitrate            int64                    `json:"MaxStreamingBitrate,omitempty"`
	StartTimeTicks                 int64                    `json:"StartTimeTicks,omitempty"`
	AudioStreamIndex               int32                    `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex            int32                    `json:"SubtitleStreamIndex,omitempty"`
	MaxAudioChannels               int32                    `json:"MaxAudioChannels,omitempty"`
	ItemId                         int64                    `json:"ItemId,omitempty"`
	DeviceProfile                  *DlnaDeviceProfile       `json:"DeviceProfile,omitempty"`
	EnableDirectPlay               bool                     `json:"EnableDirectPlay,omitempty"`
	EnableDirectStream             bool                     `json:"EnableDirectStream,omitempty"`
	EnableTranscoding              bool                     `json:"EnableTranscoding,omitempty"`
	AllowVideoStreamCopy           bool                     `json:"AllowVideoStreamCopy,omitempty"`
	AllowInterlacedVideoStreamCopy bool                     `json:"AllowInterlacedVideoStreamCopy,omitempty"`
	AllowAudioStreamCopy           bool                     `json:"AllowAudioStreamCopy,omitempty"`
	DirectPlayProtocols            []MediaInfoMediaProtocol `json:"DirectPlayProtocols,omitempty"`
}
