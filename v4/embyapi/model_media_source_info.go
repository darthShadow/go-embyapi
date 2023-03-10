/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type MediaSourceInfo struct {
	Protocol                   *MediaInfoMediaProtocol            `json:"Protocol,omitempty"`
	Id                         string                             `json:"Id,omitempty"`
	Path                       string                             `json:"Path,omitempty"`
	EncoderPath                string                             `json:"EncoderPath,omitempty"`
	EncoderProtocol            *MediaInfoMediaProtocol            `json:"EncoderProtocol,omitempty"`
	Type_                      *MediaSourceType                   `json:"Type,omitempty"`
	Container                  string                             `json:"Container,omitempty"`
	Size                       int64                              `json:"Size,omitempty"`
	Name                       string                             `json:"Name,omitempty"`
	SortName                   string                             `json:"SortName,omitempty"`
	IsRemote                   bool                               `json:"IsRemote,omitempty"`
	RunTimeTicks               int64                              `json:"RunTimeTicks,omitempty"`
	ContainerStartTimeTicks    int64                              `json:"ContainerStartTimeTicks,omitempty"`
	SupportsTranscoding        bool                               `json:"SupportsTranscoding,omitempty"`
	SupportsDirectStream       bool                               `json:"SupportsDirectStream,omitempty"`
	SupportsDirectPlay         bool                               `json:"SupportsDirectPlay,omitempty"`
	IsInfiniteStream           bool                               `json:"IsInfiniteStream,omitempty"`
	RequiresOpening            bool                               `json:"RequiresOpening,omitempty"`
	OpenToken                  string                             `json:"OpenToken,omitempty"`
	RequiresClosing            bool                               `json:"RequiresClosing,omitempty"`
	LiveStreamId               string                             `json:"LiveStreamId,omitempty"`
	BufferMs                   int32                              `json:"BufferMs,omitempty"`
	RequiresLooping            bool                               `json:"RequiresLooping,omitempty"`
	SupportsProbing            bool                               `json:"SupportsProbing,omitempty"`
	Video3DFormat              *Video3DFormat                     `json:"Video3DFormat,omitempty"`
	MediaStreams               []MediaStream                      `json:"MediaStreams,omitempty"`
	Formats                    []string                           `json:"Formats,omitempty"`
	Bitrate                    int32                              `json:"Bitrate,omitempty"`
	Timestamp                  *MediaInfoTransportStreamTimestamp `json:"Timestamp,omitempty"`
	RequiredHttpHeaders        map[string]string                  `json:"RequiredHttpHeaders,omitempty"`
	DirectStreamUrl            string                             `json:"DirectStreamUrl,omitempty"`
	TranscodingUrl             string                             `json:"TranscodingUrl,omitempty"`
	TranscodingSubProtocol     string                             `json:"TranscodingSubProtocol,omitempty"`
	TranscodingContainer       string                             `json:"TranscodingContainer,omitempty"`
	AnalyzeDurationMs          int32                              `json:"AnalyzeDurationMs,omitempty"`
	ReadAtNativeFramerate      bool                               `json:"ReadAtNativeFramerate,omitempty"`
	DefaultAudioStreamIndex    int32                              `json:"DefaultAudioStreamIndex,omitempty"`
	DefaultSubtitleStreamIndex int32                              `json:"DefaultSubtitleStreamIndex,omitempty"`
	ItemId                     string                             `json:"ItemId,omitempty"`
	ServerId                   string                             `json:"ServerId,omitempty"`
}
