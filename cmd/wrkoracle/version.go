package main

import "fmt"

/*
VersionMajor: Major version component of the current release
VersionMinor: Minor version component of the current release
VersionPatch: Patch version component of the current release
VersionMeta : Version metadata to append to the version string
*/
const (
	VersionMajor = 0
	VersionMinor = 3
	VersionPatch = 2
	VersionMeta  = "alpha"
)

// Version holds the textual version string.
var Version = func() string {
	v := fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()
