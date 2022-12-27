/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type LibraryItemLinkType string

// List of Library.ItemLinkType
const (
	ARTISTS_LibraryItemLinkType            LibraryItemLinkType = "Artists"
	ALBUM_ARTISTS_LibraryItemLinkType      LibraryItemLinkType = "AlbumArtists"
	GENRES_LibraryItemLinkType             LibraryItemLinkType = "Genres"
	STUDIOS_LibraryItemLinkType            LibraryItemLinkType = "Studios"
	TAGS_LibraryItemLinkType               LibraryItemLinkType = "Tags"
	COMPOSERS_LibraryItemLinkType          LibraryItemLinkType = "Composers"
	COLLECTIONS_LibraryItemLinkType        LibraryItemLinkType = "Collections"
	ALBUMS_LibraryItemLinkType             LibraryItemLinkType = "Albums"
	COLLECTION_FOLDERS_LibraryItemLinkType LibraryItemLinkType = "CollectionFolders"
)
