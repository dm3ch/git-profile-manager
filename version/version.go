package version

//nolint:gochecknoglobals // These variables should be set by the linker during build.
var (
	VersionNumber     = "Unknown"
	VersionCommitHash = "Unknown"
	VersionBuildDate  = "Unknown"
)
