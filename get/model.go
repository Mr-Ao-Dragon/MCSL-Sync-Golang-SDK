package get

import (
	"github.com/Mr-Ao-Dragon/MCSL-Sync-Golang-SDK/info"
)

type TaskObj struct {
	TargetCore  string
	VersionInfo info.CoreVersionInfo
	TargetPath  string
	FileName    string
}
