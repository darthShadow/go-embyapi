/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type UpdatesPackageTargetSystem string

// List of Updates.PackageTargetSystem
const (
	SERVER_UpdatesPackageTargetSystem     UpdatesPackageTargetSystem = "Server"
	MB_THEATER_UpdatesPackageTargetSystem UpdatesPackageTargetSystem = "MBTheater"
	MB_CLASSIC_UpdatesPackageTargetSystem UpdatesPackageTargetSystem = "MBClassic"
	OTHER_UpdatesPackageTargetSystem      UpdatesPackageTargetSystem = "Other"
)
