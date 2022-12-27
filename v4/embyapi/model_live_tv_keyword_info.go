/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LiveTvKeywordInfo struct {
	KeywordType *LiveTvKeywordType `json:"KeywordType,omitempty"`
	Keyword     string             `json:"Keyword,omitempty"`
}
