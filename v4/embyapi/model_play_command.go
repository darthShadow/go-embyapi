/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type PlayCommand string

// List of PlayCommand
const (
	PLAY_NOW_PlayCommand         PlayCommand = "PlayNow"
	PLAY_NEXT_PlayCommand        PlayCommand = "PlayNext"
	PLAY_LAST_PlayCommand        PlayCommand = "PlayLast"
	PLAY_INSTANT_MIX_PlayCommand PlayCommand = "PlayInstantMix"
	PLAY_SHUFFLE_PlayCommand     PlayCommand = "PlayShuffle"
)
