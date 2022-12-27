/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type RatingType string

// List of RatingType
const (
	SCORE_RatingType RatingType = "Score"
	LIKES_RatingType RatingType = "Likes"
)
