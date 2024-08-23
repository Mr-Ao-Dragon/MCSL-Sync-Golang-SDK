package util

import "github.com/hashicorp/go-version"

func JavaVersionCatch(mcVersion string) (javaVersion int) {
	java7, _ := version.NewVersion("1.7")
	java8, _ := version.NewVersion("1.12.2")
	java11, _ := version.NewVersion("1.14")
	java15, _ := version.NewVersion("1.16")
	java17, _ := version.NewVersion("1.18")
	java21, _ := version.NewVersion("1.21")
	unknownVersion, _ := version.NewVersion(mcVersion)
	switch {
	case unknownVersion.LessThanOrEqual(java7):
		return 7
	case unknownVersion.LessThanOrEqual(java8):
		return 8
	case unknownVersion.LessThanOrEqual(java11):
		return 11
	case unknownVersion.LessThanOrEqual(java15):
		return 15
	case unknownVersion.LessThanOrEqual(java17):
		return 17
	case unknownVersion.LessThanOrEqual(java21):
		return 21
	default:
		return 21
	}
}
