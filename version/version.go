package version

var (
	// These variables should be set by the linker during build

	VersionNumber     = "Unknown" //nolint:gochecknoglobals
	VersionCommitHash = "Unknown" //nolint:gochecknoglobals
	VersionBuildDate  = "Unknown" //nolint:gochecknoglobals
)
