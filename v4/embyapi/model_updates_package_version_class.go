/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UpdatesPackageVersionClass string

// List of Updates.PackageVersionClass
const (
	RELEASE_UpdatesPackageVersionClass UpdatesPackageVersionClass = "Release"
	BETA_UpdatesPackageVersionClass    UpdatesPackageVersionClass = "Beta"
	DEV_UpdatesPackageVersionClass     UpdatesPackageVersionClass = "Dev"
)
