package get

import (
	"MCSL-Sync-Golang-SDK/info"
)

type TaskObj struct {
	TargetCore  string
	VersionInfo info.CoreVersionInfo
	TargetPath  string
	FileName    string
}
